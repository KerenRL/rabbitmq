package infraestructure

import (
	"rabbitmq/src/citas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewCitasController struct {
	useCase *application.ViewCitas
}

func NewViewCitasController(useCase *application.ViewCitas) *ViewCitasController {
	return &ViewCitasController{useCase: useCase}
}

func (ec_c *ViewCitasController) Execute(c *gin.Context) {
	citas, err := ec_c.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, citas)
}