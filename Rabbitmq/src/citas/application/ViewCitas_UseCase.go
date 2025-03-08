package application

import (
	"rabbitmq/src/citas/domain"
)

type ViewCitas struct {
	db domain.ICitas
}

func NewViewCitas(db domain.ICitas) *ViewCitas {
	return &ViewCitas{db: db}
}

func (vc *ViewCitas) Execute() ([]domain.Cita, error) {
	return vc.db.GetAll()
}