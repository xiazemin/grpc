/Users/xiazemin/.gvm/pkgsets/go1.18beta2/global/pkg/mod/google.golang.org/grpc@v1.44.0/server.go


// ServiceInfo contains unary RPC method info, streaming RPC method info and metadata for a service.
type ServiceInfo struct {
	Methods []MethodInfo
	// Metadata is the metadata specified in ServiceDesc when registering service.
	Metadata interface{}
}


// GetServiceInfo returns a map from service names to ServiceInfo.
// Service names include the package names, in the form of <package>.<service>.
func (s *Server) GetServiceInfo() map[string]ServiceInfo {
     for n, srv := range s.services {
}


type Server struct {
	opts serverOptions

	mu  sync.Mutex // guards following
	lis map[net.Listener]bool
	// conns contains all active server transports. It is a map keyed on a
	// listener address with the value being the set of active transports
	// belonging to that listener.
	conns    map[string]map[transport.ServerTransport]bool
	serve    bool
	drain    bool
	cv       *sync.Cond              // signaled when connections close for GracefulStop
	services map[string]*serviceInfo // service name -> service info
	events   trace.EventLog

	quit               *grpcsync.Event
	done               *grpcsync.Event
	channelzRemoveOnce sync.Once
	serveWG            sync.WaitGroup // counts active Serve goroutines for GracefulStop

	channelzID int64 // channelz unique identification number
	czData     *channelzData

	serverWorkerChannels []chan *serverWorkerData
}


func NewServer(opt ...ServerOption) *Server {
  s := &Server{
      services: make(map[string]*serviceInfo),
  }
}

func (s *Server) register(sd *ServiceDesc, ss interface{}) {
   if _, ok := s.services[sd.ServiceName]; ok {
		logger.Fatalf("grpc: Server.RegisterService found duplicate service registration for %q", sd.ServiceName)
	}
   
   s.services[sd.ServiceName] = info
}


	info := &serviceInfo{
		serviceImpl: ss,
		methods:     make(map[string]*MethodDesc),
		streams:     make(map[string]*StreamDesc),
		mdata:       sd.Metadata,
	}



// ServiceDesc represents an RPC service's specification.
type ServiceDesc struct {
	ServiceName string
	// The pointer to the service interface. Used to check whether the user
	// provided implementation satisfies the interface requirements.
	HandlerType interface{}
	Methods     []MethodDesc
	Streams     []StreamDesc
	Metadata    interface{}
}


func (s *Server) RegisterService(sd *ServiceDesc, ss interface{}){
    s.register(sd, ss)
}


// ServiceRegistrar wraps a single method that supports service registration. It
// enables users to pass concrete types other than grpc.Server to the service
// registration methods exported by the IDL generated code.
type ServiceRegistrar interface {
	// RegisterService registers a service and its implementation to the
	// concrete type implementing this interface.  It may not be called
	// once the server has started serving.
	// desc describes the service and its methods and handlers. impl is the
	// service implementation which is passed to the method handlers.
	RegisterService(desc *ServiceDesc, impl interface{})
}

