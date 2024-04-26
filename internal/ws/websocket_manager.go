package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Tüm originlerden gelen bağlantıları kabul eder
	},
}

// WebSocketManager WebSocket bağlantılarını yönetir
type WebSocketManager struct {
	Connections map[*websocket.Conn]bool
}

// NewWebSocketManager yeni bir WebSocketManager örneği oluşturur
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Connections: make(map[*websocket.Conn]bool),
	}
}

// HandleConnection bir istemci bağlantısını kabul eder ve yönetir
func (manager *WebSocketManager) HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to websocket:", err)
		return
	}
	defer conn.Close()
	manager.Connections[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading from websocket:", err)
			break
		}
		log.Printf("Received: %s", message)
		// Burada mesajı işleme veya yönlendirme işlemleri yapabilirsiniz.
	}
	delete(manager.Connections, conn)
}

// BroadcastMessage tüm bağlı istemcilere bir mesaj gönderir
func (manager *WebSocketManager) BroadcastMessage(message string) {
	for conn := range manager.Connections {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println("Error sending message:", err)
			delete(manager.Connections, conn)
		}
	}
}
