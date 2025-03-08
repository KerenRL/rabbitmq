package main

import (
	"log"
	"rabbitmqConsumer/src/config"
	"rabbitmqConsumer/src/citas/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := infraestructure.InitCitas(); err != nil {
		log.Fatalf("Error al inicializar las Citas: %v", err)
	}
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

	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}