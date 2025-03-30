package main

import (
	"fmt"

	"crunch03/game"
	"crunch03/internal"
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
	fmt.Println(flags["verbose"])

	game.StartGame()
}
