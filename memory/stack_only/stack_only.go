package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/collectors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

func (g *Game) NextGeneration() {
	for x := range g.board {
		for y := range g.board[x] {
			neibought := g.numLiveNeiboughts(Point{x: x, y: y})
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
	g.cleanNewBoard()
}

func (g *Game) cleanNewBoard() {
	for x := range g.newBoard {
		for y := range g.newBoard[x] {
			g.newBoard[x][y] = false
		}
	}
}

func (g *Game) numLiveNeiboughts(p Point) int {
	num := 0
	if p.x > 0 && g.board[p.x-1][p.y] {
		num++
	}
	if p.x < 29 && g.board[p.x+1][p.y] {
		num++
	}
	if p.y > 0 && g.board[p.x][p.y-1] {
		num++
	}
	if p.y < 29 && g.board[p.x][p.y+1] {
		num++
	}
	if p.x > 0 && p.y > 0 && g.board[p.x-1][p.y-1] {
		num++
	}
	if p.x < 29 && p.y > 0 && g.board[p.x+1][p.y-1] {
		num++
	}
	if p.x > 0 && p.y < 29 && g.board[p.x-1][p.y+1] {
		num++
	}
	if p.y < 29 && p.x < 29 && g.board[p.x+1][p.y+1] {
		num++
	}
	return num
}

/**
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
*/

func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	life := NewGame(
		Point{x: 0, y: 1},
		Point{x: 1, y: 1},
		Point{x: 1, y: 2},
		Point{x: 2, y: 0},
		Point{x: 2, y: 1},
	)

	go func() {
		for {
			//PrintBoard(&life)
			life.NextGeneration()

			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.ListenAndServe(":8080", nil)
}
