package cmd

import (
	"github.com/jchavannes/jgo/web"
	"math/rand"
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
			game := runGame(rand.Intn(100))
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
