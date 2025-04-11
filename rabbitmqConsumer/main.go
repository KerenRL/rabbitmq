package main

import (
	"log"
	"rabbitmqConsumer/src/citas/infraestructure"
	"rabbitmqConsumer/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Inicializar las citas
	if err := infraestructure.InitCitas(); err != nil {
		log.Fatalf("Error al inicializar las Citas: %v", err)
	}

	// Inicializar repositorios
	mysqlRepo := infraestructure.NewMySQL()
	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al inicializar RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	rabbitRepo := infraestructure.NewRabbitRepository(rabbitMQRepo.Ch)

	// Configurar rutas de citas
	citasRouter := infraestructure.SetupRouter(mysqlRepo, rabbitRepo)
	for _, route := range citasRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar WebSocket
	r.GET("/ws", func(c *gin.Context) {
		config.HandleConnections(c.Writer, c.Request)
	})

	// Ejecutar manejador de mensajes en una goroutine
	go config.HandleMessages()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Iniciar servidor
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
