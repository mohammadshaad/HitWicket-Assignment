package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type GameState struct {
	Grid [5][5]string `json:"grid"`
	Turn string       `json:"turn"`
}

type Move struct {
	Player string `json:"player"`
	Move   string `json:"move"`
}

var gameState = GameState{
	Grid: [5][5]string{},
	Turn: "Player1",
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		var move Move
		if err := json.Unmarshal(msg, &move); err != nil {
			fmt.Println("Error unmarshaling message:", err)
			continue
		}

		if move.Player != gameState.Turn {
			err := conn.WriteMessage(websocket.TextMessage, []byte(`{"error":"Not your turn"}`))
			if err != nil {
				fmt.Println("Error sending message:", err)
			}
			continue
		}

		// Handle game move
		// For simplicity, we'll just toggle turns here
		gameState.Turn = "Player2"
		if gameState.Turn == "Player2" {
			gameState.Turn = "Player1"
		}

		// Broadcast updated game state
		broadcastGameState(conn)
	}
}

func broadcastGameState(conn *websocket.Conn) {
	state, err := json.Marshal(gameState)
	if err != nil {
		fmt.Println("Error marshaling game state:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, state)
	if err != nil {
		fmt.Println("Error sending game state:", err)
	}
}
