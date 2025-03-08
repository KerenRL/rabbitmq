package infraestructure

import (
	"rabbitmq/src/config"
	"log"
)

func InitCitas() error {
	log.Println("Inicializando citas...")

	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Conexión a la base de datos para citas establecida correctamente")
	return nil
}