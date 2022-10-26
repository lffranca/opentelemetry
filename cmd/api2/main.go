package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/lffranca/opentelemetry/internal/application"
	"github.com/lffranca/opentelemetry/internal/gotel"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := gotel.InitProvider()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	app := application.CreateExternalRequestApp()

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Panicln(err)
	}
}

func init() {
	requiredEnvs := []string{
		"GRPC_GO_RETRY",
		"APPLICATION_NAME",
		"PORT",
		"COLLECTOR_HOST",
		"EXTERNAL_API",
	}

	for _, env := range requiredEnvs {
		if _, ok := os.LookupEnv(env); !ok {
			log.Panicln("this environment variable is required: ", env)
		}
	}
}
