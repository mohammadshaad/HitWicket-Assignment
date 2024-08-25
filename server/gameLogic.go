package main

import (
	"strconv"
	"strings"
)

type MoveDelta struct {
	dx, dy int
}

var moveDelta = map[string]MoveDelta{
	"L":  {dx: 0, dy: -1},
	"R":  {dx: 0, dy: 1},
	"F":  {dx: -1, dy: 0},
	"B":  {dx: 1, dy: 0},
	"FL": {dx: -1, dy: -1},
	"FR": {dx: -1, dy: 1},
	"BL": {dx: 1, dy: -1},
	"BR": {dx: 1, dy: 1},
}

func isValidMove(character, move string, state *GameState) bool {
	typePos := splitCharacter(character)
	x, y := parsePosition(typePos[1])

	delta, exists := moveDelta[move]
	if !exists {
		return false
	}

	newX := x + delta.dx
	newY := y + delta.dy

	if !isValidPosition(newX, newY) {
		return false
	}

	targetCell := state.Board[newX][newY]
	if targetCell != nil && (*targetCell)[0] == character[0] {
		return false
	}

	// Handle special rules for H1 and H2
	if typePos[0] == "H1" || typePos[0] == "H2" {
		if typePos[0] == "H1" && (abs(delta.dx) > 2 || abs(delta.dy) > 2) {
			return false
		}
		if typePos[0] == "H2" && (abs(delta.dx) != abs(delta.dy) || abs(delta.dx) != 2) {
			return false
		}
	}

	return true
}

func applyMove(character, move string, state *GameState) bool {
	typePos := splitCharacter(character)
	x, y := parsePosition(typePos[1])

	delta := moveDelta[move]

	newX := x + delta.dx
	newY := y + delta.dy

	if !isValidPosition(newX, newY) {
		return false
	}

	state.Board[x][y] = nil
	state.Board[newX][newY] = &character

	return true
}

func checkWin(state *GameState) string {
	var aPieces, bPieces int
	for _, row := range state.Board {
		for _, cell := range row {
			if cell != nil {
				if (*cell)[0] == 'A' {
					aPieces++
				} else if (*cell)[0] == 'B' {
					bPieces++
				}
			}
		}
	}

	if aPieces == 0 {
		return "B"
	}
	if bPieces == 0 {
		return "A"
	}
	return ""
}

func splitCharacter(character string) []string {
	return strings.Split(character, ":")
}

func parsePosition(pos string) (int, int) {
	parts := strings.Split(pos, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return x, y
}

func isValidPosition(x, y int) bool {
	return x >= 0 && x < 5 && y >= 0 && y < 5
}

func toggleTurn(turn string) string {
	if turn == "A" {
		return "B"
	}
	return "A"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

