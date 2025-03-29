package main

import (
	"crunch03/game"
	"crunch03/internal"
	"fmt"
)

func main() {
	flags, err := internal.InitFlags()
	if err != nil {
		fmt.Println(err)
	}

	game, err := game.NewGame(flags)
	if err != nil {
		fmt.Println(err)
	}
	game.StartGame()
}
