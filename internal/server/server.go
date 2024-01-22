package server

import (
	"context"
	"fmt"
	proto "github.com/papireio/go-session-service/pkg/api/grpc"
	"google.golang.org/grpc"
	"net"
)

type clients struct {
}

type instance struct {
	proto.UnimplementedGoSessionServer
	clients *clients
}

func (i *instance) ServiceMethod(ctx context.Context, req *proto.ServiceMethodRequest) (*proto.ServiceMethodResponse, error) {
	return &proto.ServiceMethodResponse{Message: req.Message}, nil
}

func Serve(port int) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	srv := &instance{clients: &clients{}}

	grpcServer := grpc.NewServer()
	proto.RegisterGoSessionServer(grpcServer, srv)

	return grpcServer.Serve(l)
}
