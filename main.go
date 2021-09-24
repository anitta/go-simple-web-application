package main

import (
	"context"
	"log"
	"os/signal"

	"golang.org/x/sys/unix"

	"github.com/anitta/go-simple-web-application/pkg/di"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), unix.SIGINT, unix.SIGTERM, unix.SIGHUP)
	defer stop()

	err := di.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
