package main

import (
	"context"
	"github.com/isalikov/goservice/internal/env"
	"github.com/isalikov/goservice/internal/server"
	"github.com/sethvargo/go-envconfig"
	"log"
)

var ctx = context.Background()

func main() {
	config := &env.Config{}

	if err := envconfig.Process(ctx, config); err != nil {
		log.Fatalln("Fatal Error: Parsing OS ENV")
	}

	if err := server.Serve(config.Port); err != nil {
		log.Fatalln("Fatal Error: Start Up gRPC server")
	}
}
