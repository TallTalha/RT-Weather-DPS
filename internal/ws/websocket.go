package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketServer WebSocket bağlantılarını yönetir
type WebSocketServer struct {
	Upgrader websocket.Upgrader
	Clients  map[*websocket.Conn]bool
}

// NewWebSocketServer yeni bir WebSocketServer örneği oluşturur
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Tüm originlerden gelen bağlantıları kabul eder
			},
		},
		Clients: make(map[*websocket.Conn]bool),
	}
}

// HandleConnections gelen WebSocket bağlantılarını kabul eder ve işler
func (server *WebSocketServer) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := server.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading WebSocket: %v", err)
		return
	}
	defer conn.Close()

	server.Clients[conn] = true

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message: %s", p)

		// Echo the received message back
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}

// BroadcastMessage tüm bağlı istemcilere mesaj gönderir
func (server *WebSocketServer) BroadcastMessage(message []byte) {
	for conn := range server.Clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Error sending message: %v", err)
			delete(server.Clients, conn)
		}
	}
}
