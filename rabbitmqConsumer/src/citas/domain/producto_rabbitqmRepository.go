package domain

type ICitasRabbitqm interface {
	Save(order *Cita) error
}