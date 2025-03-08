package infraestructure

import (
	"rabbitmq/src/citas/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCitasController struct {
	useCase *application.CreateCitas
}

func NewCreateCitasController(useCase *application.CreateCitas) *CreateCitasController {
	return &CreateCitasController{useCase: useCase}
}

type RequestBody struct {
	Nombre string `json:"nombre"`
	Fecha    string `json:"fecha"`
	Hora     string `json:"hora"`
	Motivo   string `json:"motivo"`
}

func (cc_c *CreateCitasController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := cc_c.useCase.Execute(body.Nombre, body.Fecha, body.Hora, body.Motivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al agregar la cita", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Cita agregada correctamente"})
}