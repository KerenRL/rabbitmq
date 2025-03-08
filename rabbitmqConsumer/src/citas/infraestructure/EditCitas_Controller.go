package infraestructure

import (
	"rabbitmqConsumer/src/citas/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EditCitasController struct {
	useCase *application.EditCitas
}

func NewEditCitasController(useCase *application.EditCitas) *EditCitasController {
	return &EditCitasController{useCase: useCase}
}

func (ec_c *EditCitasController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cita inválido"})
		return
	}

	var body struct {
		Nombre string `json:"nombre"`
		Fecha string `json:"fecha"`
		Hora  string `json:"hora"`
		Motivo string `json:"motivo"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	err = ec_c.useCase.Execute(int32(id), body.Nombre, body.Fecha, body.Hora, body.Motivo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la cita"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cita actualizada correctamente"})
}