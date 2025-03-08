package infraestructure

import (
	"rabbitmq/src/config"
	"rabbitmq/src/citas/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ICitas = (*MySQL)(nil)

func NewMySQL() domain.ICitas {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveCita(nombre string, fecha string, hora string, motivo string) error {
	query := "INSERT INTO cita (nombre, fecha, hora, motivo) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, fecha, hora, motivo)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Cita guardada correctamente: Nombre:%s Fecha: %s Hora: %s - Motivo: %s", nombre, fecha, hora, motivo)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Cita, error) {
	query := "SELECT * FROM cita"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var citas []domain.Cita

	for rows.Next() {
		var cita domain.Cita
		if err := rows.Scan(&cita.ID, &cita.Nombre, &cita.Fecha, &cita.Hora, &cita.Motivo); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		citas = append(citas, cita)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
	return citas, nil
}

func (mysql *MySQL) UpdateCita(id int32, nombre string, fecha string, hora string, motivo string) error {
	query := "UPDATE cita SET nombre = ?, fecha = ?, hora = ?, motivo = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, fecha, hora, motivo, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Cita actualizada correctamente: ID: %d Nombre:%s Fecha: %s Hora: %s - Motivo: %s", id, nombre, fecha, hora, motivo)
	} else {
		log.Println("[MySQL] - No se actualizó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) DeleteCitas(id int32) error {
	query := "DELETE FROM cita WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Cita eliminada correctamente: ID: %d", id)
	} else {
		log.Println("[MySQL] - No se eliminó ninguna fila")
	}
	return nil
}