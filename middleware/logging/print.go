package logging

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		resp, err := handler(ctx, req)
		log.Println(info.Server)
		log.Println(info.FullMethod)
		log.Println(req)
		log.Println(resp)
		if err != nil {
			log.Println(err.Error())
		}

		return resp, err
	}
}
