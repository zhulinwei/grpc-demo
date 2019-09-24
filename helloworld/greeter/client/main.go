package main

import (
	"context"
	"fmt"
	pb "github.com/zhulinwei/grpc-demo/helloworld/greeter/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

type GreeterClient struct{}

const (
	port    = ":8080"
	address = "localhost:8080"
)

func (GreeterClient) Run() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	greeterClient := pb.NewGreeterClient(conn)
	result, err := greeterClient.SayHello(ctx, &pb.HelloRequest{Name: "Tony"})

	fmt.Println(result)
	if err := conn.Close(); err != nil {
		log.Fatalf("clost connect fail: %v", err)
	}
}


func main() {
	new(GreeterClient).Run()
}
