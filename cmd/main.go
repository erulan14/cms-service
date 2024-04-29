package main

import (
	"cms-service/internal/app"
	"cms-service/pkg/logging"
	"context"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("Logger initialized")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("failed to initialize app: %v", err)
	}

	err = a.Run()
	if err != nil {
		logger.Fatalf("failed to run app: %v", err)
	}
}
