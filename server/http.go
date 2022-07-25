package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-helloworld/proto/helloworld"
	"log"
	"net/http"
)

func NewHTTPServer(ctx context.Context) error {
	// Start HTTP Server
	gateway := runtime.NewServeMux()
	endpoint := fmt.Sprintf("localhost%s", GRPCPort)

	r := mux.NewRouter()
	r.PathPrefix("/").Handler(gateway)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterGreeterHandlerFromEndpoint(ctx, gateway, endpoint, opts)
	if err != nil {
		log.Fatalf("RegisterGreeterHandlerFromEndpoint: %v", err)
	}

	log.Println("Starting http server")

	return http.ListenAndServe(HTTPPort, gateway)
}
