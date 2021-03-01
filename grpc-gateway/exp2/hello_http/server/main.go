package main

import (
	"context"
	pb "grpc-gateway.http/proto/hello_http"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50151"

type server struct {
	pb.UnimplementedHelloHTTPServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloHTTPRequest) (*pb.HelloHTTPResponse, error) {
	log.Printf("Received :%v", in.GetName())
	return &pb.HelloHTTPResponse{
		Message: "hello" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloHTTPServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
