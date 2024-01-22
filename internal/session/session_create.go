package session

import (
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func Create(ctx context.Context, client *redis.Client, sessionToken string, uuid string) error {
	if sessionToken == "" || uuid == "" {
		return status.Error(codes.InvalidArgument, "Incorrect request argument")
	}

	if err := client.Set(ctx, sessionToken, uuid, time.Hour*24*30).Err(); err != nil {
		return status.Error(codes.Internal, "Internal server error (redis HSet)")
	}

	return nil
}
