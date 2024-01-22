package session

import (
	"context"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Flush(ctx context.Context, redis *redis.Client, sessionTokens []string) error {
	if len(sessionTokens) == 0 {
		return status.Error(codes.InvalidArgument, "Incorrect request argument")
	}

	pipe := redis.Pipeline()
	for _, key := range sessionTokens {
		pipe.Del(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return status.Error(codes.Internal, "Internal server error (exec redis Pipe)")
	}

	return nil
}
