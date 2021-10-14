package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/why444216978/grpc-example/proto/v1"
	resp "github.com/why444216978/grpc-example/response/v1"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	grpc_log "github.com/why444216978/grpc-example/middleware/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
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
	address := ":1234"
	Start(address, time.Second*3)
}

func Start(address string, timeout time.Duration) (int, error) {
	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_log.UnaryServerInterceptor(),
		grpc_recovery.UnaryServerInterceptor(
			grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				err = errors.WithStack(fmt.Errorf("%v", p))
				return status.Errorf(codes.Internal, "%+v", err)
			})),
	)))
	pb.RegisterHelloServiceServer(server, new(HelloService))

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return 0, err
	}

	reflection.Register(server)
	service.RegisterChannelzServiceToServer(server)

	ch := make(chan os.Signal, 1)
	go func() {
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}()

	go func() {
		fmt.Fprintf(os.Stderr, "%s server exit:%v\n", time.Now(), server.Serve(lis))
	}()

	exitCode := 0
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	select {
	case sig := <-ch:
		fmt.Fprintf(os.Stderr, "%s receive signal %v\n", time.Now(), sig)
		exitCode = 2
		time.Sleep(timeout)
		server.GracefulStop()
		fmt.Fprintf(os.Stderr, "%s exit by signal %v\n", time.Now(), sig)
	}

	return exitCode, nil
}
