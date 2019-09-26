package main

import (
	"context"
	"github.com/gin-gonic/gin"
	pb "github.com/zhulinwei/grpc-demo/helloworld/gin/proto"
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
	return &pb.HelloReply{
		ReplyValue1:  "1",
		ReplyValue2:  "2",
		ReplyValue3:  "3",
		ReplyValue4:  "4",
		ReplyValue5:  "5",
		ReplyValue6:  "6",
		ReplyValue7:  "7",
		ReplyValue8:  "8",
		ReplyValue9:  "9",
		ReplyValue10: "10",
		ReplyValue11: "11",
		ReplyValue12: "12",
		ReplyValue13: "13",
		ReplyValue14: "14",
		ReplyValue15: "15",
		ReplyValue16: "16",
		ReplyValue17: "17",
		ReplyValue18: "18",
		ReplyValue19: "19",
		ReplyValue20: "20",
	}, nil
	//return &pb.HelloReply{Message: "Hello " + req.Name}, nil
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
		//name := ctx.Query("name")
		ctx.JSON(http.StatusOK, pb.HelloReply{
			ReplyValue1:  "1",
			ReplyValue2:  "2",
			ReplyValue3:  "3",
			ReplyValue4:  "4",
			ReplyValue5:  "5",
			ReplyValue6:  "6",
			ReplyValue7:  "7",
			ReplyValue8:  "8",
			ReplyValue9:  "9",
			ReplyValue10: "10",
			ReplyValue11: "11",
			ReplyValue12: "12",
			ReplyValue13: "13",
			ReplyValue14: "14",
			ReplyValue15: "15",
			ReplyValue16: "16",
			ReplyValue17: "17",
			ReplyValue18: "18",
			ReplyValue19: "19",
			ReplyValue20: "20",
		})
		//ctx.JSON(http.StatusOK, gin.H{
		//	//"Message": "Hello " + name,
		//
		//})
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
