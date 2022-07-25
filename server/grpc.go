package server

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go-helloworld/proto/helloworld"
	"log"
	"net"
)

func NewGRPCServer() error {
	lis, err := net.Listen("tcp", GRPCPort)
	if err != nil {
		log.Fatalf("Can't listen on specified grpcPort: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	return s.Serve(lis)
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type server struct {
	pb.UnimplementedGreeterServer
}
