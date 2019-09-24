package main

import (
	"context"
	pb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Greeter struct{}

type GreeterServer struct{}

const (
	port    = ":8080"
	address = "localhost:8080"
)

// 注意需要按照greeter.proto生成后的greeter.pb.go格式传参
func (g *Greeter) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello"}, nil
}

func (GreeterServer) Run() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &Greeter{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	new(GreeterServer).Run()
}
