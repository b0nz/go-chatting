package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			if err = conn.WriteMessage(messageType, message); err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
