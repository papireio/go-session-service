package session

import (
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Delete(ctx context.Context, redis *redis.Client, sessionToken string) error {
	if sessionToken == "" {
		return status.Error(codes.InvalidArgument, "Incorrect request argument")
	}

	err := redis.Del(ctx, sessionToken)
	if err != nil {
		return status.Error(codes.Internal, "Internal server error (redis Del)")
	}

	return nil
}
