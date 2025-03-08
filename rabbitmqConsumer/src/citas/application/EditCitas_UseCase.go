package application

import (
	"rabbitmqConsumer/src/citas/domain"
)

type EditCitas struct {
	db domain.ICitas
}

func NewEditCitas(db domain.ICitas) *EditCitas {
	return &EditCitas{db: db}
}

func (ec *EditCitas) Execute(id int32, nombre string, fecha string, hora string, motivo string) error {
	return ec.db.UpdateCita(id, nombre, fecha, hora, motivo)
}