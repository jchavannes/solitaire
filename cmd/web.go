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

	resetRoute = web.Route{
		Pattern: "/reset",
		Handler: func(r *web.Response) {
			game := reset()
			r.WriteJson(game, true)
		},
	}

	fullGameRoute = web.Route{
		Pattern: "/full-game",
		Handler: func(r *web.Response) {
			fullGame := sol.FullGame{}
			fullGame.Generate(*reset())
			r.WriteJson(fullGame, true)
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
			resetRoute,
			fullGameRoute,
		},
	}
	return server.Run()
}

var game *sol.Game

func reset() *sol.Game {
	game = sol.GetGame2()
	game.FlipPiles()

	return game
}

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
