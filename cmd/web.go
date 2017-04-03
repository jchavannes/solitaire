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
			game := getNextMove()
			r.WriteJson(game, true)
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

var game *sol.Game

func getNextMove() *sol.Game {
	if game == nil {
		game = sol.GetGame2()
		game.FlipPiles()
	}

	if ! game.FindAndMakePossibleMoves() && len(game.Deck.Cards) > 0 {
		game.Deck.NextCard()
		game.Moves++
	}

	game.FlipPiles()

	return game
}
