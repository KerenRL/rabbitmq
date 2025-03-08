package domain

type ICitas interface {
	SaveCita(nombre string, fecha string, hora string, motivo string) error
	GetAll() ([]Cita, error)
	UpdateCita(id int32, nombre string, fecha string, hora string, motivo string) error
	DeleteCitas(id int32) error
}

type Cita struct {
	ID     int32  `json:"id"`
	Nombre string `json:"nombre"`
	Fecha  string `json:"fecha"`
	Hora   string `json:"hora"`
	Motivo string `json:"motivo"`
}

func NewCita(nombre string, fecha string, hora string, motivo string) *Cita {
	return &Cita{Nombre: nombre, Fecha: fecha, Hora: hora, Motivo: motivo}
}

func (c *Cita) SetNombre(nombre string) {
	c.Nombre = nombre
}