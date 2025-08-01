package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	port = ":7001"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ns, exists := md["virtual-namespace"]; exists {
			log.Printf("7001 ns is: %v", ns)
		}
	}
	log.Printf("7001 receive")

	return &pb.HelloReply{Message: "Hello " + port + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
