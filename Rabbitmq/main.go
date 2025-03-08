package main

import (
	"log"
	"rabbitmq/src/config"
	"rabbitmq/src/config/middleware"
	"rabbitmq/src/citas/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.NewCorsMiddleware())

	mysqlRepo := infraestructure.NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al inicializar RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	rabbitRepo := infraestructure.NewRabbitRepository(rabbitMQRepo.Ch)

	citasRouter := infraestructure.SetupRouter(mysqlRepo, rabbitRepo)
	for _, route := range citasRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.SetTrustedProxies([]string{"127.0.0.1"})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}