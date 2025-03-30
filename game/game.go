package game

import (
	"bufio"
	"crunch03/internal"
	"fmt"
	"time"
)

type game struct {
	config *internal.Config
	flags  map[string]interface{}
	file   string
	grid   internal.Grid
}

func NewGame(flags map[string]interface{}) (*game, error) {
	var game game
	var err error
	game.config, err = internal.InitConfig()

	if err != nil {
		return nil, err
	}

	if game.config.random {
		grid := internal.NewGrid()
		grid.InitGrid(game.config.width, game.config.height)
		grid.GenerateRandomGrid()
		game.grid = grid
	}

	if game.config.file != nil {
		scanner := bufio.NewScanner(game.config.file)
		grid := internal.NewGrid()
		fileGrid := make([][]rune, 0)
		for scanner.Scan() {
			line := scanner.Text()
			row := make([]rune, 0)
			for _, ch := range line {
				row = append(row, ch)
			}
			fileGrid = append(fileGrid, row)
		}
		grid.InitFileGrid(fileGrid)
		game.config.file.Close()
	}

	if game.config.fullscreen {
		game.grid.AdjustToTerminalSize()
	}
	game.flags = flags
	return &game, nil
}

func (g *game) CheckFlags() {
	for key, val := range g.flags {
		fmt.Print("key: " + key + " val: ")
		fmt.Println(val)
	}
}

func (g *game) StartGame() {
	for g.grid.LivingCells > 0 {

		g.grid.PrintGrid()
		time.Sleep(g.config.delay)
	}
}
