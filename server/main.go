package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type GameState struct {
	Board [5][5]*string `json:"board"`
	Turn  string        `json:"turn"`
}

var gameState = GameState{
	Board: [5][5]*string{},
	Turn:  "A",
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func broadcastGameState(clients map[*websocket.Conn]bool) {
	for client := range clients {
		if client != nil {
			err := client.WriteJSON(gameState)
			if err != nil {
				log.Printf("Error broadcasting state: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func handleConnections(clients map[*websocket.Conn]bool, ws *websocket.Conn) {
	defer ws.Close()
	clients[ws] = true
	ws.WriteJSON(gameState)

	for {
		var message map[string]interface{}
		err := ws.ReadJSON(&message)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, ws)
			break
		}

		// Remove the declaration of the 'action' variable
		data := message["data"].(map[string]interface{})
		player := data["player"].(string)
		character := data["character"].(string)
		move := data["move"].(string)

		if player != gameState.Turn {
			ws.WriteJSON(map[string]string{"error": "Not your turn"})
			continue
		}

		if !isValidMove(character, move, &gameState) {
			ws.WriteJSON(map[string]string{"error": "Invalid move"})
			continue
		}

		if !applyMove(character, move, &gameState) {
			ws.WriteJSON(map[string]string{"error": "Move failed"})
			continue
		}

		winner := checkWin(&gameState)
		if winner != "" {
			broadcastGameState(clients)
			for client := range clients {
				client.WriteJSON(map[string]string{"gameOver": winner})
			}
			return
		}

		gameState.Turn = toggleTurn(gameState.Turn)
		broadcastGameState(clients)
	}
}

func main() {
	clients := make(map[*websocket.Conn]bool)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading to websocket: %v", err)
			return
		}
		handleConnections(clients, ws)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

