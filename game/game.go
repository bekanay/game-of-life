package game

import (
	"crunch03/internal"
	"errors"
	"fmt"
	"time"
)

type game struct {
	verbose     bool
	delay       time.Duration
	edgePortals bool
	fullscreen  bool
	footprints  bool
	colored     bool
	flags       map[string]interface{}
	file        string
	grid        internal.Grid
}

func NewGame(flags map[string]interface{}) (*game, error) {
	var game game
	for key, val := range flags {
		switch key {
		case "verbose":
			if v, ok := val.(bool); ok {
				game.verbose = v
			}
		case "edges-portal":
			if v, ok := val.(bool); ok {
				game.edgePortals = v
			}
		case "fullscreen":
			if v, ok := val.(bool); ok {
				game.fullscreen = v
			}
		case "footprints":
			if v, ok := val.(bool); ok {
				game.footprints = v
			}
		case "colored":
			if v, ok := val.(bool); ok {
				game.colored = v
			}
		case "delay-ms":
			if v, ok := val.(int); ok {
				game.delay = time.Millisecond * time.Duration(v)
			}
		case "file":
			if v, ok := val.(bool); ok {
				game.colored = v
			}
		case "random":
			if values, ok := val.([]int); ok && len(values) == 2 {
				grid := internal.NewGrid()
				grid.InitGrid(values[0], values[1])
				grid.GenerateRandomGrid()
				game.grid = grid
			}
		default:
			return nil, errors.New("Warning: Unknown flag " + key)
		}
	}
	if game.fullscreen {
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
		time.Sleep(g.delay)
	}
}
