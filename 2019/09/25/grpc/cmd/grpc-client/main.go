package main

import (
	"context"
	"fmt"
	"log"

	api "github.com/lyrise/sandbox-go/2019/09/25/grpc/gens/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := api.NewAnimalClient(conn)
	message := &api.GetNameArgument{Id: 2}
	res, err := client.GetName(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
