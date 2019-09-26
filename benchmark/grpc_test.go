package benchmark

import (
	"context"
	pb "github.com/zhulinwei/grpc-demo/helloworld/gin/proto"
	"google.golang.org/grpc"
	"testing"
)

const (
	rpcAddress = "localhost:8080"
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
