package internal

import (
	"fmt"
	"math/rand"
	"os"

	"crunch03/utils"

	"golang.org/x/term"
)

var colors = []string{"\033[032m", "\033[031m", "\033[033m", "\033[0m"}

type Grid struct {
	gameMap       [][]rune
	Height        int
	Width         int
	LivingCells   int
	LivingChar    rune
	EmptyChar     rune
	FootPrintChar rune
}

func NewGrid() Grid {
	return Grid{}
}

func (g *Grid) InitDefaultCells() {
	g.LivingChar = '×'
	g.EmptyChar = '·'
	g.FootPrintChar = '∘'
}

func (g *Grid) InitCustomCells(c Config) {
	g.LivingChar = c.CustomCells[0]
	g.EmptyChar = c.CustomCells[1]
	g.FootPrintChar = c.CustomCells[2]
}

func (g *Grid) InitGrid(w, h int) {
	g.gameMap = make([][]rune, h)
	for i := 0; i < h; i++ {
		g.gameMap[i] = make([]rune, w)
	}

	g.Height = h
	g.Width = w
}

func (g *Grid) InitInputGrid(inputGrid [][]rune) {
	g.Height = len(inputGrid)
	g.Width = len(inputGrid[0])
	g.gameMap = make([][]rune, g.Height)
	for id, row := range inputGrid {
		rowGrid := make([]rune, 0)
		for _, ch := range row {
			if !(ch == '#' || ch == '.') {
				fmt.Println("Incorrect character: " + string(ch) + "\n")
				os.Exit(0)
			}
			if ch == '.' {
				rowGrid = append(rowGrid, g.EmptyChar)
			} else {
				rowGrid = append(rowGrid, g.LivingChar)
				g.LivingCells++
			}
		}
		g.gameMap[id] = rowGrid
	}
}

func (g *Grid) GenerateRandomGrid() {
	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if rand.Intn(2) == 1 {
				g.gameMap[i][j] = g.LivingChar
				g.LivingCells++
			} else {
				g.gameMap[i][j] = g.EmptyChar
			}
		}
	}
}

func (g *Grid) AdjustToTerminalSize(config *Config) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	width = width / 2
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}
	if g.Height < height {
		for g.Height < height {
			g.AddRow()
		}
	} else if g.Height > height {
		g.gameMap = g.gameMap[:height]
		g.Height = height
	}

	if g.Width < width {
		for g.Width < width {
			g.AddColumn()
		}
	} else if g.Width > width {
		for i := range g.gameMap {
			g.gameMap[i] = g.gameMap[i][:width]
		}
		g.Width = width
	}
	if config.Fullscreen && config.Verbose && g.Height > 8 {
		g.Height -= 5
	}
}

func (g *Grid) AddRow() {
	newRow := make([]rune, g.Width)
	for i := 0; i < g.Width; i++ {
		newRow[i] = g.EmptyChar
	}
	g.gameMap = append(g.gameMap, newRow)
	g.Height++
}

func (g *Grid) AddColumn() {
	for i := 0; i < g.Height; i++ {
		g.gameMap[i] = append(g.gameMap[i], g.EmptyChar)
	}
	g.Width++
}

func (g *Grid) PrintGrid(config *Config, tick int) {
	clearScreen()
	if config.Verbose {
		fmt.Printf("Tick: %d\n", tick)
		fmt.Printf("Grid Size: %dx%d\n", g.Width, g.Height)
		fmt.Printf("Live Cells: %d\n", g.LivingCells)
		fmt.Printf("DelayMs: %dms\n\n", config.Delay.Abs().Milliseconds())
	}

	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if config.Colored {
				if g.gameMap[i][j] == g.LivingChar {
					fmt.Print(colors[0])
				} else if g.gameMap[i][j] == g.EmptyChar {
					fmt.Print(colors[1])
				} else {
					fmt.Print(colors[2])
				}
				fmt.Print(string(g.gameMap[i][j]))
				fmt.Print(colors[3])
			} else {
				fmt.Print(string(g.gameMap[i][j]))
			}
			if j != g.Width-1 {
				fmt.Print(" ")
			}
		}
		if i != g.Height-1 {
			fmt.Println()
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (g *Grid) UpdateGird(config *Config) {
	newGrid := make([][]rune, g.Height)

	for i := 0; i < g.Height; i++ {
		newGrid[i] = make([]rune, g.Width)
	}
	g.LivingCells = 0

	for i := 0; i < g.Height; i++ {
		for j := 0; j < g.Width; j++ {
			if utils.IsAlive(g.gameMap, i, j, g.LivingChar, config.EdgePortals) {
				newGrid[i][j] = g.LivingChar
				g.LivingCells++
			} else {
				if g.gameMap[i][j] == g.LivingChar || g.gameMap[i][j] == g.FootPrintChar {
					newGrid[i][j] = g.FootPrintChar
				} else {
					newGrid[i][j] = g.EmptyChar
				}
			}
		}
	}

	g.gameMap = newGrid
}
