package main

import (
	"log"

	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/config"
	_ "github.com/masioware/mercado-livre-desafio-tecnico/order-api/docs"
	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/internal/router"
	"github.com/masioware/mercado-livre-desafio-tecnico/order-api/pkg/logger"
)

func main() {
	logger.Init()

	config.LoadEnv()
	config.InitMongo()

	app := router.NewRouter()

	if err := app.Run(":" + config.GetPort()); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
