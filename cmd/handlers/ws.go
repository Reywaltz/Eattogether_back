package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

func HandleConnections(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer conn.Close()

	clients[conn] = true

	fmt.Printf("User Connected: %v\n", conn.RemoteAddr())

	for {
		var msg Message
		_, msg_byte, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
		}

		err = json.Unmarshal(msg_byte, &msg)
		if err != nil {
			fmt.Printf("Can't unmarshall to msg %v\n", err)
		}

		broadcast <- msg
	}
}

func HandleBroadcast() {
	for {
		message := <-broadcast
		for client := range clients {
			byte_message, _ := json.Marshal(message)
			fmt.Printf("Message from broadcast: %v\n", string(byte_message))
			err := client.WriteMessage(websocket.TextMessage, []byte("Hello from server"))
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
