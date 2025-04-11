package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // Lista de clientes conectados
var broadcast = make(chan string)            // Canal para enviar mensajes

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleConnections maneja las conexiones WebSocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error al actualizar a WebSocket: %v", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Cliente desconectado: %v", err)
			delete(clients, ws)
			break
		}
	}
}

// HandleMessages envía mensajes a todos los clientes conectados
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("Error al enviar mensaje: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

// NotifyNewCita envía una notificación cuando se agrega una nueva cita
func NotifyNewCita(citaID string) {
	message := "Nueva cita agregada con ID: " + citaID
	broadcast <- message
}

// NotifyNewCitaDetails envía una notificación con los detalles de la cita
func NotifyNewCitaDetails(nombre, fecha, hora, motivo string) {
	message := fmt.Sprintf("Nueva cita agregada: Nombre: %s, Fecha: %s, Hora: %s, Motivo: %s", nombre, fecha, hora, motivo)
	broadcast <- message
}
