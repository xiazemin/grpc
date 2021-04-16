实现这个解析器。

type mydnsResolver struct {
	domain       string
	port         string
	address      map[resolver.Address]struct{}
	ctx    context.Context
	cancel context.CancelFunc
	cc     resolver.ClientConn
	wg sync.WaitGroup
}
 
// ResolveNow resolves immediately
func (mr *mydnsResolver) ResolveNow(resolver.ResolveNowOptions) {
	select {
	case mr.rn <- struct{}{}:
	default:
	}
}
 
// Close stops resolving
func (mr *mydnsResolver) Close() {
	mr.cancel()
	mr.wg.Wait()
}
 
func (mr *mydnsResolver) watcher() {
	defer util.CheckPanic()
	defer mr.wg.Done()
 
	for {
		select {
		case <-mr.ctx.Done():
			return
		case <-mr.rn:
		}
		result, err := mr.resolveByHttpDNS()
		if err != nil || len(result) == 0 {
			continue
		}
		mr.cc.UpdateState(resolver.State{Addresses: result})
	}
}
 
func (mr *mydnsResolver) resolveByHttpDNS() ([]resolver.Address, error) {
	var items []string = make([]string, 0, 4)
 
	//这里实现通过向http://myself.dns.xyz发送get请求获取实例ip列表，并存入items中
 
	var addresses = make([]resolver.Address, 0, len(items))
	for _, v := range items {
		addr := net.JoinHostPort(v, mr.port)
		a := resolver.Address{
			Addr:       addr,
			ServerName: addr, // same as addr
			Type:       resolver.Backend,
		}
		addresses = append(addresses, a)
	}
 
	return addresses, nil

}


type mydnsBuilder struct {
}
 
func NewBuilder() resolver.Builder {
	return &mydnsBuilder{}
}
 
// Scheme for mydns
func (mb *mydnsBuilder) Scheme() string {
	return "mydns"
}
 
// Build
func (mb *mydnsBuilder) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	host, port, err := net.SplitHostPort(target.Endpoint)
	if err != nil {
		host = target.Endpoint
		port = "80"
	}
 
	ctx, cancel := context.WithCancel(context.Background())
	mr := &mydnsResolver{
		domain:       host,
		port:         port,
		cc:           cc,
		rn:           make(chan struct{}, 1),
		address:      make(map[resolver.Address]struct{}),
	}
	mr.ctx, pr.cancel = ctx, cancel
 
	mr.wg.Add(1)
	go mr.watcher()
 
	mr.ResolveNow(resolver.ResolveNowOptions{})
	return mr, nil
}
接下来我们还要实现，当这个包初始化时，将scheme注册到grpc的解析器map中：

func init() {
	resolver.Register(NewBuilder())
}
实现好这个包之后，我们只需要在调用Dial的文件import mydns这个包，并且保证传入的target满足以下格式：

mydns://service.order
grpc就会使用我们实现的解析器，向我们自己的dns服务器请求服务对应的ip地址了




