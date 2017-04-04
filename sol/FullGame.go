package sol

type FullGame struct {
	Piles [7]Pile
	Deck  Deck
	Moves []Move
	Won   bool
}

func (g *FullGame) Generate(game Game) {
	g.Piles = game.Piles
	g.Deck = game.Deck

	for i := 0; i < 500; i++ {
		moves := game.FindPossibleMoves()
		g.Moves = append(g.Moves, moves...)
		for _, move := range moves {
			game.MakeMove(move)
		}
		if len(moves) == 0 {
			if len(game.Deck.Cards) == 0 {
				break
			}
			game.Deck.NextCard()
			game.Moves++
		} else {
			game.FlipPiles()
		}
	}

}
