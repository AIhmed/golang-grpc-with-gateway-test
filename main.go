package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc-test/proto/testdata" // Updated import
	"grpc-test/server"
)

func main() {
	// Start gRPC server in a goroutine
	go func() {
		server.StartGRPCServer()
	}()

	// Create a client connection to the gRPC server
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	// Create a new ServeMux for the gRPC-Gateway
	gwmux := runtime.NewServeMux()

	// Register TestDataService
	err = testdata.RegisterTestDataServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("HTTP server started on port 8080")
	log.Fatalln(gwServer.ListenAndServe())
}
