package middlewares

import (
	"context"
	"time"

	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"google.golang.org/grpc"
)

func WithServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(serverInterceptor)
}

func serverInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	// Calls the handler
	var h interface{}
	var err error
	h, err = handler(ctx, req)

	config.GrpcLog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
