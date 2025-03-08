package application

import (
	"rabbitmq/src/citas/domain"
)

type CreateCitas struct {
	rabbit domain.ICitasRabbitqm
	db domain.ICitas
}

func NewCreateCitas(rabbit domain.ICitasRabbitqm, db domain.ICitas) *CreateCitas {
	return &CreateCitas{rabbit: rabbit, db: db}
}

func (cc *CreateCitas) Execute(nombre string, fecha string, hora string, motivo string) error {
	err := cc.db.SaveCita(nombre, fecha, hora, motivo)
	if err != nil {
		return err
	}

	cita := domain.NewCita(nombre, fecha, hora, motivo)

	err = cc.rabbit.Save(cita)
	if err != nil {
		return err		
	}

	return nil
}
