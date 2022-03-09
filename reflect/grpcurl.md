grpcurl.go

dial := func() *grpc.ClientConn {
    cc, err := grpcurl.BlockingDial(ctx, network, target, creds, opts...)
}
		
refClient = grpcreflect.NewClient(refCtx, reflectpb.NewServerReflectionClient(cc))
reflSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)


svcs, err := grpcurl.ListServices(descSource)
methods, err := grpcurl.ListMethods(descSource, symbol)
svcs, err := descSource.ListServices()
txt, err := grpcurl.GetDescriptorText(dsc, descSource)





desc_source.go
func DescriptorSourceFromServer(_ context.Context, refClient *grpcreflect.Client) DescriptorSource {
	return serverSource{client: refClient}
}

type DescriptorSource interface {
	// ListServices returns a list of fully-qualified service names. It will be all services in a set of
	// descriptor files or the set of all services exposed by a gRPC server.
	ListServices() ([]string, error)
	// FindSymbol returns a descriptor for the given fully-qualified symbol name.
	FindSymbol(fullyQualifiedName string) (desc.Descriptor, error)
	// AllExtensionsForType returns all known extension fields that extend the given message type name.
	AllExtensionsForType(typeName string) ([]*desc.FieldDescriptor, error)
}


type serverSource struct {
	client *grpcreflect.Client
}

func (ss serverSource) ListServices() ([]string, error) {
	svcs, err := ss.client.ListServices()
	return svcs, reflectionSupport(err)
}


github.com/jhump/protoreflect@v1.10.3/grpcreflect/client.go
func (cr *Client) ListServices() ([]string, error) {
	req := &rpb.ServerReflectionRequest{
		MessageRequest: &rpb.ServerReflectionRequest_ListServices{
			// proto doesn't indicate any purpose for this value and server impl
			// doesn't actually use it...
			ListServices: "*",
		},
	}
	resp, err := cr.send(req)
	if err != nil {
		return nil, err
	}

	listResp := resp.GetListServicesResponse()
	if listResp == nil {
		return nil, &ProtocolError{reflect.TypeOf(listResp).Elem()}
	}
	serviceNames := make([]string, len(listResp.Service))
	for i, s := range listResp.Service {
		serviceNames[i] = s.Name
	}
	return serviceNames, nil
}











describe


dsc, err := descSource.FindSymbol(s)

fqn := dsc.GetFullyQualifiedName()
var elementType string
switch d := dsc.(type) {
case *desc.MessageDescriptor:
    elementType = "a message"
    parent, ok := d.GetParent().(*desc.MessageDescriptor)



func (ss serverSource) FindSymbol(fullyQualifiedName string) (desc.Descriptor, error){
    	file, err := ss.client.FileContainingSymbol(fullyQualifiedName)
        d := file.FindSymbol(fullyQualifiedName)
}


client.go
func (cr *Client) FileContainingSymbol(symbol string) (*desc.FileDescriptor, error) {
    	req := &rpb.ServerReflectionRequest{
		MessageRequest: &rpb.ServerReflectionRequest_FileContainingSymbol{
			FileContainingSymbol: symbol,
		},
	}
}
