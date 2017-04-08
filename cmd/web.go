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
			game := sol.GetGame1()
			game.FlipPiles()
			fullGame := sol.FullGame{}
			fullGame.Generate(*game)
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
		},
	}
	return server.Run()
}
