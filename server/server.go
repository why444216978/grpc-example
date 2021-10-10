package main

import (
	"context"
	"log"
	"net"

	pb "github.com/why444216978/grpc-example/proto/v1"
	resp "github.com/why444216978/grpc-example/response/v1"

	"google.golang.org/grpc"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (p *HelloService) Hello(ctx context.Context, req *pb.Request) (*resp.Response, error) {
	data := resp.ResponseData{
		Name: req.GetName(),
		Age:  18,
	}

	reply := &resp.Response{
		Code:    0,
		Message: "",
		Data:    &data,
	}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
