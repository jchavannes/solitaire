package sol

type FullGame struct {
	Piles       [7]Pile
	Deck        Deck
	Moves       []Move
	NoPileCards []NoPileCard
	Won         bool
}

type NoPileCard struct {
	Card  Card
	Times int
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
		if game.Deck.Position == 0 && len(game.Deck.Cards) > 0 {
			game.Deck.NextCard()
			game.Moves++
			g.Moves = append(g.Moves, Move{
				SourcePileId: PileDeck,
				TargetPileId: PileDeck,
				SourceCard: game.Deck.Cards[0],
			})
		}
		moves := game.FindPossibleMoves()
		for _, move := range moves {
			if game.MakeMove(move) {
				g.Moves = append(g.Moves, move)
			}
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
		if game.IsGameCompleted() {
			g.Won = true
			break
		}
	}

}

func (g *FullGame) Optimize() {
	// Loop through all moves
	for i := 0; i < len(g.Moves); i++ {
		move := g.Moves[i]
		// Ignore moves that are from deck or top of a pile
		if move.SourcePileId == PileDeck || move.SourcePileIndex == 0 {
			continue
		}
		// Loop backwards through moves, starting from current move being checked
		for h := i - 1; h > 0; h-- {
			compareMove := g.Moves[h]
			// If current move target pile equals compare move source pile
			if move.TargetPileId == compareMove.SourcePileId || move.TargetPileId == compareMove.TargetPileId {
				// If same card being moved, remove card
				if compareMove.SourceCard == move.SourceCard && move.SourcePileId == compareMove.TargetPileId {
					//fmt.Println("Removing moves:")
					// Loop through all moves in between
					for j := h + 1; j < i; j++ {
						checkMove := &g.Moves[j]
						// If moves have same source pile, update check move to other new source pile
						if checkMove.SourcePileId == move.SourcePileId {
							//fmt.Printf("Updating move: %#v to use source pile id: %d\n", checkMove, compareMove.SourcePileId)
							checkMove.SourcePileId = compareMove.SourcePileId
							checkMove.SourcePileIndex += (compareMove.SourcePileIndex - move.SourcePileIndex)
						}
						// If check move target is current move source, update to other new move source
						if checkMove.TargetPileId == move.SourcePileId {
							//fmt.Printf("Updating move: %#v to use target pile id: %d\n", checkMove, move.TargetPileId)
							checkMove.TargetPileId = move.TargetPileId
						}
					}
					//fmt.Printf("CompareMove: %#v\nMove: %#v\n", compareMove, move)
					//fmt.Printf("i: %d, h: %d, move.TargetPileId: %d, compareMove.SourcePileId: %d, move.SourceCard: %#v, compareMove.SourceCard: %#v\n", i, h, move.TargetPileId, compareMove.SourcePileId, move.SourceCard, compareMove.SourceCard)
					g.Moves = append(g.Moves[:i], g.Moves[i + 1:]...)
					g.Moves = append(g.Moves[:h], g.Moves[h + 1:]...)
					i--
				}
				break
			}
		}
	}
}
