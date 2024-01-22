package session

import (
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Extract(ctx context.Context, redis *redis.Client, sessionToken string) (string, error) {
	if sessionToken == "" {
		return "", status.Error(codes.InvalidArgument, "Incorrect request argument")
	}

	uuid, err := redis.Get(ctx, sessionToken).Result()
	if err != nil {
		return "", status.Error(codes.NotFound, "Session not exist or expired")
	}

	return uuid, nil
}
