type GRPCServer interface {
	grpc.ServiceRegistrar
	GetServiceInfo() map[string]grpc.ServiceInfo
}


type serverReflectionServer struct {
	rpb.UnimplementedServerReflectionServer
	s GRPCServer

	initSymbols  sync.Once
	serviceNames []string
	symbols      map[string]*dpb.FileDescriptorProto // map of fully-qualified names to files
}

func Register(s GRPCServer) {
	rpb.RegisterServerReflectionServer(s, &serverReflectionServer{
		s: s,
	})
}


func (s *serverReflectionServer) ServerReflectionInfo(stream rpb.ServerReflection_ServerReflectionInfoServer) error {
   		switch req := in.MessageRequest.(type) {
		case *rpb.ServerReflectionRequest_FileByFilename:
        
        case *rpb.ServerReflectionRequest_FileContainingSymbol:
			b, err := s.fileDescEncodingContainingSymbol(req.FileContainingSymbol, sentFileDescriptors)
      
       case *rpb.ServerReflectionRequest_FileContainingExtension:

       case *rpb.ServerReflectionRequest_AllExtensionNumbersOfType:

       case *rpb.ServerReflectionRequest_ListServices:
          svcNames, _ := s.getSymbols()
}


func (s *serverReflectionServer) getSymbols() (svcNames []string, symbolIndex map[string]*dpb.FileDescriptorProto) 

  serviceInfo := s.s.GetServiceInfo()

  		for svc, info := range serviceInfo {
			s.serviceNames = append(s.serviceNames, svc)
			fdenc, ok := parseMetadata(info.Metadata)

            fd, err := decodeFileDesc(fdenc)
            s.processFile(fd, processed)



func (s *serverReflectionServer) fileDescEncodingByFilename(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
    enc := proto.FileDescriptor(name)
    return fileDescWithDependencies(fd, sentFileDescriptors)
}

// allExtensionNumbersForTypeName returns all extension numbers for the given type.
func (s *serverReflectionServer) allExtensionNumbersForTypeName(name string) ([]int32, error) {
    st, err := typeForName(name)
    extNums, err := s.allExtensionNumbersForType(st)
}






case *rpb.ServerReflectionRequest_FileContainingSymbol:
b, err := s.fileDescEncodingContainingSymbol(req.FileContainingSymbol, sentFileDescriptors)





func (s *serverReflectionServer) fileDescEncodingContainingSymbol(name string, sentFileDescriptors map[string]bool) ([][]byte, error) {
    fd, err = s.fileDescForType(st)
}

func (s *serverReflectionServer) fileDescForType(st reflect.Type) (*dpb.FileDescriptorProto, error) {
	m, ok := reflect.Zero(reflect.PtrTo(st)).Interface().(protoMessage)
    enc, _ := m.Descriptor()
	return decodeFileDesc(enc)
}

func decodeFileDesc(enc []byte) (*dpb.FileDescriptorProto, error) {
	raw, err := decompress(enc)
    fd := new(dpb.FileDescriptorProto)
    if err := proto.Unmarshal(raw, fd); err != nil {
}

func decompress(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
}



// protoMessage is used for type assertion on proto messages.
// Generated proto message implements function Descriptor(), but Descriptor()
// is not part of interface proto.Message. This interface is needed to
// call Descriptor().
type protoMessage interface {
	Descriptor() ([]byte, []int)
}




func (*ServerReflectionRequest) Descriptor() ([]byte, []int) {
	return file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDescGZIP(), []int{0}
}


func file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDescGZIP() []byte {
		file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDescData)
}


	file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDescData = file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDesc




var file_reflection_grpc_reflection_v1alpha_reflection_proto_rawDesc = []byte{
	0x0a, 0x33, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70,
}


/Users/xiazemin/.gvm/pkgsets/go1.18beta2/global/pkg/mod/github.com/golang/protobuf@v1.5.2/protoc-gen-go/descriptor/descriptor.pb.go

type FileDescriptorProto = descriptorpb.FileDescriptorProto

