package cmd

import (
	"github.com/jchavannes/jgo/web"
	"github.com/jchavannes/solitaire/sol"
)

var (
	indexRoute = web.Route{
		Pattern: "/",
		Handler: func(r *web.Response) {
			r.Render()
		},
	}

	gameRoute = web.Route{
		Pattern: "/game",
		Handler: func(r *web.Response) {
			var leastMoves sol.FullGame
			preGame := getFullGame([]sol.NoPileCard{})
			for i := 0; i < 24; i++ {
				for j := 0; j < 24; j++ {
					card := preGame.Deck.Cards[i]
					secondCard := preGame.Deck.Cards[j]
					fullGame := getFullGame([]sol.NoPileCard{
						{
							Card: card,
							Times: 1000,
						},
						{
							Card: secondCard,
							Times: 1000,
						},
					})
					lessMoves := len(fullGame.Moves) < len(leastMoves.Moves)
					winAlreadyAndThisIsALoss := leastMoves.Won == true && fullGame.Won == false
					firstWin := leastMoves.Won == false && fullGame.Won == true
					better := !winAlreadyAndThisIsALoss && (lessMoves || firstWin)
					if len(leastMoves.Moves) == 0 || better {
						leastMoves = fullGame
					}
				}
			}
			r.WriteJson(leastMoves, true)
		},
	}
)

func runWeb() error {
	server := web.Server{
		Port: 8250,
		StaticFilesDir: "web",
		TemplatesDir: "web",
		Routes: []web.Route{
			indexRoute,
			gameRoute,
		},
	}
	return server.Run()
}

func getFullGame(noPileCards []sol.NoPileCard) sol.FullGame {
	game := sol.GetGame7()
	game.FlipPiles()
	game.NoPileCards = noPileCards
	fullGame := sol.FullGame{}
	fullGame.Generate(*game)
	fullGame.Optimize()
	fullGame.NoPileCards = noPileCards
	return fullGame
}
