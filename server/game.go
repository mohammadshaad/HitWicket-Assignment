// game.go
package main

import (
	"fmt"
)

// Character represents a game character
type Character struct {
	Type  string
	Owner string
	X     int
	Y     int
}

// Game represents the game state
type Game struct {
	Board    [][]*Character
	Players  [2]string
	Turn     int
	GameOver bool
}

// Initialize a new game
func NewGame(player1, player2 string) *Game {
	game := &Game{
		Board:   make([][]*Character, 5),
		Players: [2]string{player1, player2},
		Turn:    0,
	}

	// Initialize the game board (5x5 grid)
	for i := range game.Board {
		game.Board[i] = make([]*Character, 5)
	}

	// Place initial characters (this could be modified as needed)
	game.Board[0][0] = &Character{Type: "Pawn", Owner: player1, X: 0, Y: 0}
	game.Board[0][1] = &Character{Type: "Hero1", Owner: player1, X: 0, Y: 1}
	game.Board[4][3] = &Character{Type: "Hero2", Owner: player2, X: 4, Y: 3}
	game.Board[4][4] = &Character{Type: "Pawn", Owner: player2, X: 4, Y: 4}

	return game
}

// MakeMove processes a move command
func (g *Game) MakeMove(player, move string) bool {
	// Logic for handling moves (parsing move string, checking validity, updating game state)
	// Example: "P1:L" -> Move Pawn owned by player 1 to the left
	fmt.Printf("Player %s made move: %s\n", player, move)

	// Return true if the move was valid, otherwise false
	return true
}
