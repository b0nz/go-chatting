package main

import (
	"fmt"
	"net/http"
)

// type M map[string]interface{}

// const MESSAGE_NEW_USER = "New User"
// const MESSAGE_CHAT = "Chat"
// const MESSAGE_LEAVE = "Leave"

// var connections = make([]*WebSocketConnection, 0)

// type SocketPayload struct {
// 	Message string
// }

// type SocketResponse struct {
// 	From    string
// 	Type    string
// 	Message string
// }

// type WebSocketConnection struct {
// 	*websocket.Conn
// 	Username string
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// websocket
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
