package game

import (
	"bufio"
	"fmt"
	"time"

	"crunch03/internal"
)

type game struct {
	config *internal.Config
	flags  map[string]interface{}
	grid   internal.Grid
}

func NewGame(flags map[string]interface{}) (*game, error) {
	var game game
	var err error
	game.config, err = internal.InitConfig(flags)
	if err != nil {
		return nil, err
	}

	if game.config.Random {
		grid := internal.NewGrid()
		grid.InitGrid(game.config.Width, game.config.Height)
		grid.GenerateRandomGrid()
		game.grid = grid
	}

	if game.config.File != nil {
		scanner := bufio.NewScanner(game.config.File)
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
		game.config.File.Close()
	}

	if game.config.Fullscreen {
		game.grid.AdjustToTerminalSize(game.config)
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
	tick := 1
	g.grid.PrintGrid(g.config, tick)
	for g.grid.LivingCells > 0 {
		g.grid.UpdateGird(g.config)
		g.grid.PrintGrid(g.config, tick)
		tick++
		time.Sleep(g.config.Delay)
	}
}
