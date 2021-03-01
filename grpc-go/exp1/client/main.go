package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "exp1/google.golang.org/grpc/examples/helloworld/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50151"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet :%v", err)
	}

	log.Printf("Greeting:%s", r.GetMessage())
}
