package main

import (
	"log"
	"net"

	api "github.com/lyrise/sandbox-go/2019/09/25/grpc/gens/api"
	service "github.com/lyrise/sandbox-go/2019/09/25/grpc/rpc/service"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	animalService := &service.AnimalService{}
	api.RegisterAnimalServer(server, animalService)
	server.Serve(listenPort)
}
