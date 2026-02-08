package main

import (
	"fmt"

	"net/http"
	_ "net/http/pprof"
)

type Point struct {
	x, y int
}

const BOARD_SIZE = 30

func NewGame(startPoints ...Point) Game {
	g := Game{}
	for _, p := range startPoints {
		g.board[p.x][p.y] = true
	}
	return g
}

type Game struct {
	board    [BOARD_SIZE][BOARD_SIZE]bool
	newBoard [BOARD_SIZE][BOARD_SIZE]bool
}

func NextGeneration(g *Game) {
	for x := range g.board {
		for y := range g.board[x] {
			neibought := newLiveNeighbors(g, Point{x: x, y: y})
			if g.board[x][y] && (neibought == 2 || neibought == 3) {
				g.newBoard[x][y] = true
			} else if !g.board[x][y] && neibought == 3 {
				g.newBoard[x][y] = true
			} else {
				g.newBoard[x][y] = false
			}
		}
	}
	g.board = g.newBoard
	cleanNewBoard(g)
}

func cleanNewBoard(g *Game) {
	for x := range g.newBoard {
		for y := range g.newBoard[x] {
			g.newBoard[x][y] = false
		}
	}
}

func newLiveNeighbors(g *Game, p Point) int {
	num := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			x := (p.x + dx + BOARD_SIZE) % BOARD_SIZE
			y := (p.y + dy + BOARD_SIZE) % BOARD_SIZE
			if g.board[x][y] {
				num++
			}
		}
	}
	return num
}

func PrintBoard(g *Game) {
	fmt.Print("\033[H")
	for x := range g.board {
		for y := range g.board[x] {
			if g.board[x][y] {
				fmt.Print("X")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	life := NewGame(
		Point{x: 0, y: 1},
		Point{x: 1, y: 1},
		Point{x: 1, y: 2},
		Point{x: 2, y: 0},
		Point{x: 2, y: 1},
	)

	for range 1000000 {
		NextGeneration(&life)
	}
}
