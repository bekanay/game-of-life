package internal

import (
	"fmt"
	"math/rand"
)

type Grid struct {
	gameMap     [][]rune
	height      int
	width       int
	LivingCells int
	LivingChar  rune
	EmptyChar   rune
}

func NewGrid() Grid {
	return Grid{}
}

func (g *Grid) InitGrid(w, h int) {
	g.gameMap = make([][]rune, h)
	for i := 0; i < h; i++ {
		g.gameMap[i] = make([]rune, w)
	}
	g.LivingChar = '×'
	g.EmptyChar = '·'
	g.height = h
	g.width = w
}

func (g *Grid) GenerateRandomGrid() {
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			if rand.Intn(2) == 1 {
				g.gameMap[i][j] = g.LivingChar
				g.LivingCells++
			} else {
				g.gameMap[i][j] = g.EmptyChar
			}
		}
	}
}

func (g *Grid) PrintGrid() {
	fmt.Print("\033[H\033[2J") // clear terminal
	fmt.Print("\033[?25l")     // remove cursor
	for i := 0; i < g.height; i++ {
		for j := 0; j < g.width; j++ {
			fmt.Print(string(g.gameMap[i][j]))
			if j != g.width-1 {
				fmt.Print(" ")
			}
		}
		if i != g.height-1 {
			fmt.Println()
		}
	}
}
