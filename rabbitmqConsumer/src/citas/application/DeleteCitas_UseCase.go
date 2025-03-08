package application

import (
	"rabbitmqConsumer/src/citas/domain"
)

type DeleteCitas struct {
	db domain.ICitas
}

func NewDeleteCitas(db domain.ICitas) *DeleteCitas {	
	return &DeleteCitas{db: db}
}

func (dc *DeleteCitas) Execute(id int32) error {
	return dc.db.DeleteCitas(id)
}