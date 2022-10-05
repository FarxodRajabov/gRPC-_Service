package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"todo-service/connect"
	"todo-service/internal"
	"todo-service/proto"
)

func main() {
	log.Println("Starting listening on port 8082")
	port := ":8082"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s", port)
	conn := connect.Connection()
	srv := internal.NewGRPCServer(conn)

	grpcServer := grpc.NewServer()
	proto.RegisterTodoServiceServer(grpcServer, srv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
