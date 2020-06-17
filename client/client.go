package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-go-helloworld/helloworld/grpc/helloworld"
	"log"
	"time"
)

const address = "localhost:50080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can't connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Kamal"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Greetings: %s", res.GetMessage())
}
