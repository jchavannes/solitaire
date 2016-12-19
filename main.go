package main

import (
	"github.com/jchavannes/solitaire/sol"
)

func main() {
	println("Solitaire.")
	game := sol.GetSampleGame()
	game.OutputMissingCards()

	game.FlipPiles()

	for i := 0; i < 150; i++ {
		if ! game.FindAndMakePossibleMoves() {
			if len(game.Deck.Cards) == 0 {
				break
			}
			game.Deck.NextCard()
			game.Moves++
		} else {
			game.FlipPiles()
		}
		if i % 10 == 0 {
			//game.OutputGame()
		}
	}
	game.OutputGame()

	//fmt.Printf("%#v\n", game.FindPossibleMoves())
}
