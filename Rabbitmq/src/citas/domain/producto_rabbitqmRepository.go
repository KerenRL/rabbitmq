package domain

type ICitasRabbitqm interface {
	Save(cita *Cita) error
}