package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/why444216978/grpc-example/proto/v1"
	resp "github.com/why444216978/grpc-example/response/v1"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	grpc_log "github.com/why444216978/grpc-example/middleware/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
	// panic(111)
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_log.UnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(
			grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				err = errors.WithStack(fmt.Errorf("%v", p))
				return status.Errorf(codes.Internal, "%+v", err)
			})),
	)))
	pb.RegisterHelloServiceServer(grpcServer, new(HelloService))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
		return
	}

	reflection.Register(grpcServer)

	grpcServer.Serve(lis)
}
