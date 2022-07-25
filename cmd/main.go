package main

import (
	"context"
	"grpc-go-helloworld/server"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println("Starting http server")
	go server.NewHTTPServer(ctx)
	log.Println("Starting grpc server")
	server.NewGRPCServer()
}
