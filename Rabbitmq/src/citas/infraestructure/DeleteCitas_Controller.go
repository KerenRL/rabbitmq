package infraestructure

import (
	"rabbitmq/src/citas/application"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCitasController struct {
	useCase *application.DeleteCitas
}

func NewDeleteCitasController(useCase *application.DeleteCitas) *DeleteCitasController {
	return &DeleteCitasController{useCase: useCase}
}

func (dc_c *DeleteCitasController) Execute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cita inválido"})
		return
	}

	err = dc_c.useCase.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error al eliminar la cita: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cita eliminada correctamente"})
}