package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	pb "github.com/zhulinwei/grpc-demo/helloworld/gin/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	port       = ":3000"
	rpcAddress = "localhost:8080"
	ginAddress = "localhost:8081"
)

type Router struct{}

func (Router) InitGRPCRoute(route *gin.Engine) {
	var err error
	var conn *grpc.ClientConn

	if conn, err = grpc.Dial(rpcAddress, grpc.WithInsecure()); err != nil {
		log.Fatal("did not connect: %v", err)
	}
	greeterClient := pb.NewGreeterClient(conn)

	route.GET("/api/grpc/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		req := &pb.HelloRequest{Name: name}

		var result *pb.HelloReply
		if result, err = greeterClient.SayHello(ctx, req); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(),})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(result.Message),})
	})
}

func (Router) InitHTTPRoute(route *gin.Engine) {
	route.GET("/api/http/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		uri := url.URL{
			Scheme:   "http",
			Host:     ginAddress,
			Path:     "/http",
			RawQuery: "name=" + name,
		}
		var err error
		var response *http.Response
		if response, err = http.Get(uri.String()); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if response == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": "nothing"})
			return
		}
		var body []byte
		if body, err = ioutil.ReadAll(response.Body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err = response.Body.Close(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		type Body struct {
			Message string
		}
		var obj Body
		err = json.Unmarshal(body, &obj)
		ctx.JSON(http.StatusOK, gin.H{
			"result": obj.Message,
		})
	})
}

func main() {
	route := gin.Default()
	router := new(Router)
	router.InitGRPCRoute(route)
	router.InitHTTPRoute(route)
	if err := route.Run(port); err != nil {
		log.Fatal("gin run fail: %v", err)
	}
}
