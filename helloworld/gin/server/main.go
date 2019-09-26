package main

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Greeter struct{}

type GreeterServer struct{}

const (
	rpcPort = ":8080"
	ginPort = ":8081"
)

// 注意需要按照greeter.proto生成后的greeter.pb.go格式传参
func (g *Greeter) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func (GreeterServer) RunGRPC(port string) {
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

func (GreeterServer) RunHTTP(port string) {
	route := gin.Default()
	route.GET("/http", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Hello " + name,
		})
	})
	if err := route.Run(port); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	greeterServer := new(GreeterServer)
	go greeterServer.RunGRPC(rpcPort)
	greeterServer.RunHTTP(ginPort)
}
