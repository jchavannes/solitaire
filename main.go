package main

import (
	"github.com/jchavannes/solitaire/sol"
)

func main() {
	println("Solitaire.")
	game := sol.GetSampleGame()
	game.OutputMissingCards()

	game.FlipPiles()
	game.OutputGame()

	for i := 0; i < 10; i++ {
		if ! game.FindAndMakePossibleMoves() {
			println("No new moves.")
			break
		}
		game.FlipPiles()
		game.OutputGame()
	}

	//fmt.Printf("%#v\n", game)
}
