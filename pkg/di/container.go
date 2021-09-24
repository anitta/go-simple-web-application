package di

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/anitta/go-simple-web-application/pkg/infrastructure"
	"github.com/anitta/go-simple-web-application/pkg/infrastructure/config"
	"github.com/anitta/go-simple-web-application/pkg/interfaces/controllers"
)

type DI struct{}

func Start(ctx context.Context) error {
	di := DI{}

	env, err := config.Get()
	if err != nil {
		return err
	}

	router := infrastructure.Router(di.SimpleController(), env.CORSAllowOrigins)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func (di *DI) SimpleController() controllers.SimpleController {
	return controllers.NewSimpleController()
}
