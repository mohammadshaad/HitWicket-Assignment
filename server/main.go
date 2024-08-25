// main.go
package main

import (
	// "encoding/json" 
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity
	},
}

type Message struct {
	Player string `json:"player"`
	Move   string `json:"move"`
}

var game *Game

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	// Send initial game state
	initialState := map[string]string{"message": "Game started"}
	err = ws.WriteJSON(initialState)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			break
		}

		fmt.Printf("Received move: %+v\n", msg)

		if game.MakeMove(msg.Player, msg.Move) {
			err = ws.WriteJSON(map[string]string{"status": "move accepted"})
		} else {
			err = ws.WriteJSON(map[string]string{"status": "invalid move"})
		}
		if err != nil {
			fmt.Println("Error writing JSON:", err)
			break
		}
	}
}

func main() {
	// Initialize the game
	game = NewGame("Player1", "Player2")

	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
