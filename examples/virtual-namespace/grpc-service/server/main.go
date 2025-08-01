package main

import (
	"context"
	"log"
	"net"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/examples/virtual-namespace/grpc-service/virtual"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	port    = ":50051"
	address = "localhost:50007"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

var greeterClient pb.GreeterClient

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ns, exists := md["virtual-namespace"]; exists {
			log.Printf("ns is: %v", ns)
		}
	}

	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "hello")

	r2, err := greeterClient.SayHello(ctx, &pb.HelloRequest{Name: "Dapr"})
	if err != nil {
		log.Printf("could not greet: %v", err)
	}
	log.Println("r", r2)

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// create the client
	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(virtual.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(virtual.StreamClientInterceptor()))
	if err != nil {
		panic(err)
	}
	client := dapr.NewClientWithConnection(conn)
	defer client.Close()
	greeterClient = pb.NewGreeterClient(client.GrpcClientConn())

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	daprServer := daprd.NewServiceWithGrpcServer(lis, s)

	// start the server
	if err := daprServer.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
