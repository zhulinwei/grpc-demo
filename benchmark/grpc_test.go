package benchmark

import (
	"context"
	"encoding/json"
	pb "github.com/zhulinwei/grpc-demo/helloworld/gin/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

const (
	rpcAddress = "localhost:8080"
	ginAddress = "localhost:8081"
	//rpcAddress = "34.217.41.67:8082"
)

func BenchmarkGRPCP(b *testing.B) {
	var err error
	var conn *grpc.ClientConn

	if conn, err = grpc.Dial(rpcAddress, grpc.WithInsecure()); err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}

	greeterClient := pb.NewGreeterClient(conn)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		var err error
		var response *pb.HelloReply
		if response, err = greeterClient.SayHello(context.Background(), &pb.HelloRequest{Name: "tony"}); err != nil {
			b.Fatalf("grpc request failed: %v", err)
		}
		if response == nil {
			b.Fatalf("grpc response is wrong: %v", response)
		}
	}
}

func BenchmarkHTTP(b *testing.B) {
	uri := url.URL{
		Scheme:   "http",
		Host:     ginAddress,
		Path:     "/http",
		RawQuery: "name=tony",
	}
	var err error
	var response *http.Response

	for n := 0; n < b.N; n++ {
		if response, err = http.Get(uri.String()); err != nil {
			b.Fatalf("http request failed: %v", err)
		}
		if response == nil {
			b.Fatalf("http response is wrong: %v", response)
		}

		var body []byte
		if body, err = ioutil.ReadAll(response.Body); err != nil {
			b.Fatalf("http body is wrong: %v", err)
		}
		if err = response.Body.Close(); err != nil {
			b.Fatalf("http body close fail: %v", err)
		}

		type Body struct {
			//Message string
			replyValue1  string
			replyValue2  string
			replyValue3  string
			replyValue4  string
			replyValue5  string
			replyValue6  string
			replyValue7  string
			replyValue8  string
			replyValue9  string
			replyValue10 string
			replyValue11 string
			replyValue12 string
			replyValue13 string
			replyValue14 string
			replyValue15 string
			replyValue16 string
			replyValue17 string
			replyValue18 string
			replyValue19 string
			replyValue20 string
		}
		var obj Body
		err = json.Unmarshal(body, &obj)
	}
	//fmt.Println(string(body))
}
