package main

import (
	"github.com/jchavannes/solitaire/sol"
)

func main() {
	println("Solitaire.")
	game := sol.GetSampleGame()
	game.OutputMissingCards()

	game.FlipPiles()

	for i := 0; i < 120; i++ {
		if ! game.FindAndMakePossibleMoves() {
			game.Deck.NextCard()
			game.Moves++
			continue
		}
		game.FlipPiles()
	}
	game.OutputGame()

	//fmt.Printf("%#v\n", game)

}