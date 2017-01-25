package main

import (
	"github.com/jchavannes/solitaire/sol"
	"math/rand"
	"fmt"
)

func main() {
	println("Solitaire.")
	var bestGames []*sol.Game
	for i := 0; i < 10000; i++ {
		game := runGame(rand.Intn(100))
		//game.OutputGame()
		if len(bestGames) < 5 {
			bestGames = append(bestGames, game)
		} else {
			for j, bestGame := range bestGames {
				if ! game.IsGameCompleted() && bestGame.IsGameCompleted() {
					continue
				}
				if game.IsGameCompleted() && ! bestGame.IsGameCompleted() {
					bestGames[j] = game
					break
				} else if game.Moves < bestGame.Moves {
					fmt.Printf("i: %d, j: %d, game.Moves: %d, game.IsGameCompleted: %t\n", i, j, game.Moves, game.IsGameCompleted())
					bestGames[j] = game
					break
				}
			}
		}
		//game.OutputGameSimple()
	}
	println()

	for _, game := range bestGames {
		game.OutputGameSimple()
	}

	//fmt.Printf("%#v\n", game.FindPossibleMoves())
}

func runGame(skipPercent int) *sol.Game {
	game := sol.GetGame1()
	game.OutputMissingCards()
	game.SetSkipPercent(skipPercent)
	game.FlipPiles()

	for i := 0; i < 500; i++ {
		if ! game.FindAndMakePossibleMoves() {
			if len(game.Deck.Cards) == 0 {
				break
			}
			game.Deck.NextCard()
			game.Moves++
		} else {
			game.FlipPiles()
		}
		if i % 25 == 0 {
			//game.OutputGame()
		}
	}
	return game
}
