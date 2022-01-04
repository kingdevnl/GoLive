package GoLive

import (
	"github.com/gofiber/websocket/v2"
	"log"
	"sync"
)

type IWebSocket interface {
	IsConnected() bool
	ID() string
	Send(message interface{}) error
	Receive() (interface{}, error)
	Close() error
}

type WebSocket struct {
	ID          string
	IsConnected bool
	Conn        *websocket.Conn
	Data        Map
	lock        sync.Mutex
}

// Send sends a message to the client.
func (ws *WebSocket) Send(json Map) {

	if !ws.IsConnected {
		return
	}
	ws.lock.Lock()
	err := ws.Conn.WriteJSON(json)

	if err != nil {
		log.Println("Error sending json to WebSocket: ", err)
	}

	ws.lock.Unlock()
}

// Receive receives a message from the client and returns a bool and the message. If the connection is closed or the message could not be received, the bool will be false.
func (ws *WebSocket) Receive() (bool, Map) {

	if !ws.IsConnected {
		return false, nil
	}

	var json Map
	err := ws.Conn.ReadJSON(&json)

	if err != nil {
		log.Println("Error receiving json from WebSocket: ", err)
		return false, nil
	}

	return true, json
}

// Close closes the connection with the client.
func (ws *WebSocket) Close() error {
	return ws.Conn.Close()
}

// NewWebsocket creates a new websocket instance.
func NewWebsocket(conn *websocket.Conn) *WebSocket {
	return &WebSocket{
		ID:          GenerateID(id_len),
		IsConnected: true,
		Conn:        conn,
		Data:        make(Map),
	}
}
