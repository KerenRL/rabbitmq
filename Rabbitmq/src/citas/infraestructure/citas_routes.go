package infraestructure

import (
	"rabbitmq/src/citas/application"
	"rabbitmq/src/citas/domain"
	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.ICitas, rabbitRepo domain.ICitasRabbitqm) *gin.Engine {
	r := gin.Default()

	createCita := application.NewCreateCitas(rabbitRepo, repo)
	createCitaController := NewCreateCitasController(createCita)

	viewCitas := application.NewViewCitas(repo)
	viewCitasController := NewViewCitasController(viewCitas)

	editCitaUseCase := application.NewEditCitas(repo)
	editCitaController := NewEditCitasController(editCitaUseCase)

	deleteCitaUseCase := application.NewDeleteCitas(repo)
	deleteCitaController := NewDeleteCitasController(deleteCitaUseCase)

	r.POST("/citas", createCitaController.Execute)
	r.GET("/citas", viewCitasController.Execute)
	r.PUT("/citas/:id", editCitaController.Execute)
	r.DELETE("/citas/:id", deleteCitaController.Execute)

	return r
}