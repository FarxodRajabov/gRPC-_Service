package main

import (
	"api-gateway/api"
	"api-gateway/services"
)

func main() {

	grpcClients, _ := services.NewGrpcClients()

	server := api.NewServer(grpcClients)

	err := server.Run(":8081")

	if err != nil {
		return
	}
}
