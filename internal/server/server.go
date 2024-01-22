package server

import (
	"context"
	"fmt"
	"github.com/papireio/go-session-service/internal/session"
	proto "github.com/papireio/go-session-service/pkg/api/grpc"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"net"
)

type Clients struct {
	RedisClient *redis.Client
}

type instance struct {
	proto.UnimplementedGoSessionServer
	clients *Clients
}

func (i *instance) CreateSession(ctx context.Context, req *proto.CreateSessionRequest) (*proto.CreateSessionResponse, error) {
	if err := session.Create(ctx, i.clients.RedisClient, req.SessionToken, req.Uuid); err != nil {
		return &proto.CreateSessionResponse{
			Success: false,
		}, err
	}

	return &proto.CreateSessionResponse{
		Success: true,
	}, nil
}

func (i *instance) ExtractSession(ctx context.Context, req *proto.ExtractSessionRequest) (*proto.ExtractSessionResponse, error) {
	uuid, err := session.Extract(ctx, i.clients.RedisClient, req.SessionToken)
	if err != nil {
		return &proto.ExtractSessionResponse{
			Uuid:    "",
			Success: false,
		}, err
	}

	return &proto.ExtractSessionResponse{
		Uuid:    uuid,
		Success: true,
	}, nil
}

func (i *instance) DeleteSession(ctx context.Context, req *proto.DeleteSessionRequest) (*proto.DeleteSessionResponse, error) {
	err := session.Delete(ctx, i.clients.RedisClient, req.SessionToken)
	if err != nil {
		return &proto.DeleteSessionResponse{
			Success: false,
		}, err
	}

	return &proto.DeleteSessionResponse{
		Success: true,
	}, nil
}

func (i *instance) FlushSessions(ctx context.Context, req *proto.FlushSessionsRequest) (*proto.FlushSessionsResponse, error) {
	err := session.Flush(ctx, i.clients.RedisClient, req.SessionTokens)
	if err != nil {
		return &proto.FlushSessionsResponse{
			Success: false,
		}, err
	}

	return &proto.FlushSessionsResponse{
		Success: true,
	}, nil
}

func Serve(port int, clients *Clients) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	srv := &instance{clients: clients}

	grpcServer := grpc.NewServer()
	proto.RegisterGoSessionServer(grpcServer, srv)

	return grpcServer.Serve(l)
}
