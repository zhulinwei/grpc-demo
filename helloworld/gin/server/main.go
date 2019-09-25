package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Greeter struct{}

type GreeterServer struct{}

const (
	rpcPort    = ":8080"
	ginPort    = ":8081"
)

// 注意需要按照greeter.proto生成后的greeter.pb.go格式传参
func (g *Greeter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func (GreeterServer) Run(port string) {
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
	go func() {
		new(GreeterServer).Run(rpcPort)
	}()
	if err := gin.Default().Run(ginPort); err != nil {
		fmt.Println(err)
	}
}
