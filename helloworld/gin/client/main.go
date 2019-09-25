package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/zhulinwei/grpc-demo/helloworld/gin/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"time"
)

const (
	port = ":3000"
	rpcAddress = "localhost:8080"
)

func main () {
	// Set up a http server.
	route := gin.Default()
	route.GET("/api/grpc/:name", func(content *gin.Context) {
		conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
		if err != nil {
			log.Fatal("did not connect: %v", err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		name := content.Param("name")
		req := &pb.HelloRequest{Name: name}
		fmt.Println(name)

		greeterClient := pb.NewGreeterClient(conn)
		result, err := greeterClient.SayHello(ctx, req)
		fmt.Println(result)
		if err != nil {
			content.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		content.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(result.Message),
		})
		if err := conn.Close(); err != nil {
			log.Fatal("close connect fail: %v", err)
		}
	})

	if err := route.Run(port); err != nil {
		log.Fatal("gin run fail: %v", err)
	}
}