package main

import (
	"context"
	pb "exp1/google.golang.org/grpc/examples/helloworld/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":50151"

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received :%v", in.GetName())
	return &pb.HelloReply{
		Message: "hello" + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
