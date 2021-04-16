如何根据服务名称，返回实例的ip？
这个问题有很多种解决方案，我们可以使用一些成熟的服务发现组件，例如consul或者zookeeper，也可以我们自己实现一个解析服务器；

如何将我们选择的服务解析方式应用到grpc的连接建立中去？
grpc的resolver，就是帮我们解决这个问题的


1.程序启动时，客户端是如何从一个域名/服务名，获取到其对应的实例ip，然后与之建立连接的呢？
2.运行过程中，如果后端的实例挂了，grpc如何感知到，并重新建立连接呢？

使用grpc的时候，首先要做的就是调用Dial或DialContext函数来初始化一个clientConn对象，而resolver是这个连接对象的一个重要的成员，所以我们首先看一看clientConn对象创建过程中，resolver是怎么设置进去的。

cc := &ClientConn{
		target:            target,
		csMgr:             &connectivityStateManager{},
		conns:             make(map[*addrConn]struct{}),
		dopts:             defaultDialOptions(),
		blockingpicker:    newPickerWrapper(),
		czData:            new(channelzData),
		firstResolveEvent: grpcsync.NewEvent(),
	}

cc.parsedTarget = grpcutil.ParseTarget(cc.target)

func ParseTarget(target string) (ret resolver.Target) {
	var ok bool
	ret.Scheme, ret.Endpoint, ok = split2(target, "://")
	if !ok {
		return resolver.Target{Endpoint: target}
	}
	ret.Authority, ret.Endpoint, ok = split2(ret.Endpoint, "/")
	if !ok {
		return resolver.Target{Endpoint: target}
	}
	return ret
}

resolverBuilder := cc.getResolver(cc.parsedTarget.Scheme)

func (cc *ClientConn) getResolver(scheme string) resolver.Builder {
	for _, rb := range cc.dopts.resolvers {
		if scheme == rb.Scheme() {
			return rb
		}
	}
	return resolver.Get(scheme)
}

其实就是在根据scheme进行查找，如果resolver已经在调用DialContext的时候通过opts参数传了进来，那我们就直接用，否则调用resolver.Get(scheme)去找，我们项目中就是用的resolver.Get(scheme)，所以我们再来看看这里是怎么做的

func Get(scheme string) Builder {
	if b, ok := m[scheme]; ok {
		return b
	}
	return nil
}

Get函数是通过m这个map，去查找有没有scheme对应的resolver的builder，那么m这个map是什么时候插入的值呢？这个在resolver的Register函数里：
func Register(b Builder) {
	m[b.Scheme()] = b
}

那么谁会去调用这个Register函数，向map中写入resolver呢？

grpc实现了一个默认的解析器，也就是"passthrough"，这个看名字就理解了，就是透传，所谓透传就是，什么都不做，那么什么时候需要透传呢？当你调用DialContext的时候，如果传入的target本身就是一个ip+port，这个时候，自然就不需要再解析什么了。那么"passthrough"对应的这个默认的解析器是什么时候注册到m这个map中的呢？这个调用在passthrough包的init函数里

func init() {
	resolver.Register(&passthroughBuilder{})
}


而grpc包会import这个_ "google.golang.org/grpc/internal/resolver/passthrough"包，所以，"passthrough"这个解析器在grpc初始化的时候就会被注册到m这个map中。


服务名称的解析会根据我们项目使用的命名系统的不同，而存在多种不同的方案，如果我们是使用consul实现服务发现，那么我们就希望我们的解析器实现的是通过concul客户端获取服务信息，而如果我们用的是dns服务，那么我们的解析器就应该通过我们的dns服务器去获取指定域名对应的服务器ip


resolverBuilder := cc.getResolver(cc.parsedTarget.Scheme)
	if resolverBuilder == nil {
		// If resolver builder is still nil, the parsed target's scheme is
		// not registered. Fallback to default resolver and set Endpoint to
		// the original target.
		channelz.Infof(cc.channelzID, "scheme %q not registered, fallback to default scheme", cc.parsedTarget.Scheme)
		cc.parsedTarget = resolver.Target{
			Scheme:   resolver.GetDefaultScheme(),
			Endpoint: target,
		}
		resolverBuilder = cc.getResolver(cc.parsedTarget.Scheme)
		if resolverBuilder == nil {
			return nil, fmt.Errorf("could not get resolver for default scheme: %q", cc.parsedTarget.Scheme)
		}
	}

scheme会被设置为默认的scheme，这个默认的scheme又是啥呢？

defaultScheme = "passthrough"


DialContext函数会使用获取到的resolver的builder，构建一个resolver，并将其赋给cc这个对象：

    // Build the resolver.
	rWrapper, err := newCCResolverWrapper(cc, resolverBuilder)
	if err != nil {
		return nil, fmt.Errorf("failed to build resolver: %v", err)
	}
	cc.mu.Lock()
	cc.resolverWrapper = rWrapper
	cc.mu.Unlock()

func newCCResolverWrapper(cc *ClientConn, rb resolver.Builder) (*ccResolverWrapper, error) {
	ccr := &ccResolverWrapper{
		cc:   cc,
		done: grpcsync.NewEvent(),
	}
 
	var credsClone credentials.TransportCredentials
	if creds := cc.dopts.copts.TransportCredentials; creds != nil {
		credsClone = creds.Clone()
	}
	rbo := resolver.BuildOptions{
		DisableServiceConfig: cc.dopts.disableServiceConfig,
		DialCreds:            credsClone,
		CredsBundle:          cc.dopts.copts.CredsBundle,
		Dialer:               cc.dopts.copts.Dialer,
	}
 
	var err error
	// We need to hold the lock here while we assign to the ccr.resolver field
	// to guard against a data race caused by the following code path,
	// rb.Build-->ccr.ReportError-->ccr.poll-->ccr.resolveNow, would end up
	// accessing ccr.resolver which is being assigned here.
	ccr.resolverMu.Lock()
	defer ccr.resolverMu.Unlock()
	ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)
	if err != nil {
		return nil, err
	}
	return ccr, nil
}
这个函数最重要的一行，就是调用了我们传入的builder的Build方法，也就是这一行：

ccr.resolver, err = rb.Build(cc.parsedTarget, ccr, rbo)

Dial函数只是创建了一个抽象的连接，实际的http2连接并没有在这里创建

底层http2连接对应的是一个grpc的stream，而stream的创建有两种方式，一种就是我们主动去创建一个stream池，这样当有请求需要发送时，我们可以直接使用我们创建好的stream


func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

创建过程中的调用过程：

newClientStream
newAttemptLocked
getTransport 
pick
getReadyTransport
connect
resetTransport
tryAllAddrs
createTransport
NewClientTransport
newHTTP2Client

https://blog.csdn.net/u013536232/article/details/108556544
https://blog.csdn.net/weixin_39526564/article/details/111363726
https://www.sohu.com/a/368368405_657921

https://www.cnblogs.com/haima/p/12239543.html

https://github.com/grpc/grpc-go/blob/master/resolver/resolver.go
https://github.com/grpc/grpc-go/blob/master/resolver_conn_wrapper.go
https://github.com/ilisin/grpc
http://morecrazy.github.io/2018/08/14/grpc-go%E5%9F%BA%E4%BA%8Eetcd%E5%AE%9E%E7%8E%B0%E6%9C%8D%E5%8A%A1%E5%8F%91%E7%8E%B0%E6%9C%BA%E5%88%B6/

https://www.cnblogs.com/zhangboyu/p/7452725.html


https://segmentfault.com/a/1190000018424798?utm_source=tag-newest

https://blog.csdn.net/qq_35916684/article/details/104055246

