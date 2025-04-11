package application

import (
	"rabbitmqConsumer/src/citas/domain"
	"rabbitmqConsumer/src/config"
)

type CreateCitas struct {
	rabbit domain.ICitasRabbitqm
	db     domain.ICitas
}

func NewCreateCitas(rabbit domain.ICitasRabbitqm, db domain.ICitas) *CreateCitas {
	return &CreateCitas{rabbit: rabbit, db: db}
}

func (cc *CreateCitas) Execute(nombre string, fecha string, hora string, motivo string) error {
	// Guardar la cita en la base de datos
	err := cc.db.SaveCita(nombre, fecha, hora, motivo)
	if err != nil {
		return err
	}

	// Crear una nueva instancia de la cita
	cita := domain.NewCita(nombre, fecha, hora, motivo)

	// Enviar la cita a RabbitMQ
	err = cc.rabbit.Save(cita)
	if err != nil {
		return err
	}

	// Notificar a través del WebSocket con los detalles de la cita
	config.NotifyNewCitaDetails(nombre, fecha, hora, motivo)

	return nil
}
