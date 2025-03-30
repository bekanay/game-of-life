package game

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
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

	if len(game.config.CustomCells) == 0 {
		game.grid.InitDefaultCells()
	} else {
		game.grid.InitCustomCells(*game.config)
	}

	if game.config.Random {
		game.grid.InitGrid(game.config.Width, game.config.Height)
		game.grid.GenerateRandomGrid()
	}

	if game.config.File != nil {
		err = InitFileGrid(&game)
		if err != nil {
			return nil, err
		}
	}

	if game.config.Height == 0 && game.config.Width == 0 {
		err := UserInputGrid(&game)
		if err != nil {
			return nil, err
		}
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
		time.Sleep(g.config.Delay)
		g.grid.UpdateGird(g.config)
		tick++
		g.grid.PrintGrid(g.config, tick)

	}
	fmt.Println()
}

func InitFileGrid(g *game) error {
	scanner := bufio.NewScanner(g.config.File)

	fileGrid := make([][]rune, 0)
	scanner.Scan()
	firstLine := scanner.Text()

	var height, width int
	var err error
	num := ""
	for _, ch := range firstLine {
		if ch == ' ' && num != "" {
			height, err = strconv.Atoi(num)
			if err != nil {
				return err
			}
			if height < 3 {
				return errors.New("specified value:" + strconv.Itoa(height) + " is too low")
			}

			num = ""
			continue
		}
		if !(ch >= 48 && ch <= 57) {
			return errors.New("incorrect height or width")
		}
		num += string(ch)
	}
	width, err = strconv.Atoi(num)
	if err != nil {
		return err
	}
	if width < 3 {
		return errors.New("specified value:" + strconv.Itoa(width) + " is too low")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != width {
			return errors.New("invalid number of characters in line")
		}
		row := make([]rune, 0)
		for _, ch := range line {
			if !(ch == '.' || ch == '#') {
				return errors.New("Entered incorrect symbol: " + string(ch))
			}
			row = append(row, ch)
		}
		fileGrid = append(fileGrid, row)
		height--
	}
	if height != 0 {
		return errors.New("invalid number of lines")
	}
	g.grid.InitInputGrid(fileGrid)
	g.config.File.Close()
	g.config.Height = height
	g.config.Width = width
	return nil
}

func UserInputGrid(g *game) error {
	var height, width int
	n, err := fmt.Scanf("%d %d\n", &height, &width)
	if err != nil || n != 2 {
		return err
	}
	if height < 3 || width < 3 {
		return errors.New("invalid size of the grid")
	}
	inputGrid := make([][]rune, 0)

	for range height {
		line := ""
		n, err := fmt.Scanf("%s\n", &line)
		if n != 1 || err != nil {
			return err
		}

		if len(line) != width {
			return errors.New("Invalid size of the line, expected: " + strconv.Itoa(width))
		}

		row := make([]rune, 0)
		for _, ch := range line {
			if !(ch == '.' || ch == '#') {
				return errors.New("Entered incorrect symbol: " + string(ch))
			}
			row = append(row, ch)
		}
		inputGrid = append(inputGrid, row)
	}

	g.grid.InitInputGrid(inputGrid)
	return nil
}
