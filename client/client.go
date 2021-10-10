package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc/proto/v1"

	"google.golang.org/grpc"
)

func main() {
	cc, err := newClientConn("localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloServiceClient(cc)

	reply, err := client.Hello(context.Background(), &pb.Request{Name: "why"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetData())
}

func newClientConn(target string) (*grpc.ClientConn, error) {
	cc, err := grpc.Dial(
		target,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	return cc, nil
}
