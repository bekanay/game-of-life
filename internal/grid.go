package internal

import (
	"fmt"
	"math/rand"
	"os"

	"golang.org/x/term"
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

func (g *Grid) AdjustToTerminalSize() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	width = width/2 + 1
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}

	if g.height < height {
		for g.height < height {
			g.AddRow()
		}
	} else if g.height > height {
		g.gameMap = g.gameMap[:height]
		g.height = height
	}

	if g.width < width {
		for g.width < width {
			g.AddColumn()
		}
	} else if g.width > width {
		for i := range g.gameMap {
			g.gameMap[i] = g.gameMap[i][:width]
		}
		g.width = width
	}
}

func (g *Grid) AddRow() {
	newRow := make([]rune, g.width)
	for i := 0; i < g.width; i++ {
		newRow[i] = g.EmptyChar
	}
	g.gameMap = append(g.gameMap, newRow)
	g.height++
}

func (g *Grid) AddColumn() {
	for i := 0; i < g.height; i++ {
		g.gameMap[i] = append(g.gameMap[i], g.EmptyChar)
	}
	g.width++
}

func (g *Grid) PrintGrid() {
	clearScreen()

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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
