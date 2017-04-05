package sol

type FullGame struct {
	Piles [7]Pile
	Deck  Deck
	Moves []Move
	Won   bool
}

func (g *FullGame) Generate(game Game) {
	for i, pile := range game.Piles {
		for _, card := range pile.StackCards {
			g.Piles[i].StackCards = append(g.Piles[i].StackCards, card)
		}
		for _, card := range pile.BaseCards {
			g.Piles[i].BaseCards = append(g.Piles[i].BaseCards, card)
		}
	}
	for _, card := range game.Deck.Cards {
		g.Deck.Cards = append(g.Deck.Cards, card)
	}

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
			g.Moves = append(g.Moves, Move{
				SourcePileId: PileDeck,
				TargetPileId: PileDeck,
				SourceCard: game.Deck.Cards[game.Deck.Position - 1],
			})
			game.Deck.NextCard()
			game.Moves++
		} else {
			game.FlipPiles()
		}
	}

}
