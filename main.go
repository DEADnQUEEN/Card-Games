package main

import (
	"cardGames/foolGame"
)

func main() {
	var game, err = foolGame.NewFoolGame(36, 1)
	if err != nil {
		panic(err)
	}

	game.ShowAllPlayableCards()
}
