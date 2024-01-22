package main

import (
	"context"
	"github.com/papireio/go-session-service/internal/env"
	"github.com/papireio/go-session-service/internal/server"
	"github.com/redis/go-redis/v9"
	"github.com/sethvargo/go-envconfig"
	"log"
)

var ctx = context.Background()

func main() {
	config := &env.Config{}

	if err := envconfig.Process(ctx, config); err != nil {
		log.Fatalln("Fatal Error: Parsing OS ENV")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisURL,
	})

	clients := &server.Clients{RedisClient: redisClient}

	if err := server.Serve(config.Port, clients); err != nil {
		log.Fatalln("Fatal Error: Start Up gRPC server")
	}
}
