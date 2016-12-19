package sol

import (
	"fmt"
)

type Game struct {
	Piles       [7]Pile
	Foundations [4]Foundation
	Deck        Deck
	Moves       int
}

func (g *Game) FlipPiles() bool {
	flipped := false
	for pile := 0; pile < 7; pile++ {
		if g.Piles[pile].Flip() {
			flipped = true
		}
	}
	return flipped
}

func (g *Game) OutputMissingCards() {
	for _, suit := range []Suit{Hearts, Diamonds, Spades, Clubs} {
		for number := 1; number <= 13; number++ {
			found := 0
			for pile := 0; pile < 7; pile++ {
				if IsCardInPile(g.Piles[pile].BaseCards, Card{Number: number, Suit: suit}) {
					found++
				}
			}
			if IsCardInPile(g.Deck.Cards, Card{Number: number, Suit: suit}) {
				found++
			}
			if found > 1 {
				fmt.Printf("Found duplicate card (Number: %d, Suit: %s)\n", number, suit)
			}
			if found < 1 {
				fmt.Printf("Unable to find card (Number: %d, Suit: %s)\n", number, suit)
			}
		}
	}
}

func (g *Game) OutputGame() {
	found := true
	printString := ""
	for pile := 0; pile < 7; pile++ {
		if len(g.Piles[pile].BaseCards) == 0 {
			printString = printString + "  -  "
		} else {
			printString = printString + fmt.Sprintf("  %1d  ", len(g.Piles[pile].BaseCards))
		}
	}
	println(printString)
	for cardNum := 0; found; cardNum++ {
		found = false
		printString = ""
		for pile := 0; pile < 7; pile++ {
			if len(g.Piles[pile].StackCards) > cardNum {
				card := g.Piles[pile].StackCards[cardNum]
				printString = printString + fmt.Sprintf(" %3s ", card.GetString())
				found = true
			} else {
				printString = printString + "     "
			}
		}
		println(printString)
	}
	currentCard, err := g.Deck.GetCurrentCard()
	if err != nil {
		println("Deck empty")
	} else {
		fmt.Printf("Deck: %d - %s\n", g.Deck.Position, currentCard.GetString())
	}
}

func (g *Game) FindPossibleMoves() []Move {
	var possibleMoves []Move
	for sourcePileId := 0; sourcePileId < 7; sourcePileId++ {
		sourcePile := g.Piles[sourcePileId]
		for targetPileId := 0; targetPileId < 7; targetPileId++ {
			targetPile := g.Piles[targetPileId]
			if targetPileId == sourcePileId || len(sourcePile.StackCards) == 0 || ! targetPile.CanMoveCardToPile(sourcePile.StackCards[0]) {
				continue
			}
			if sourcePile.StackCards[0].Number == 13 && len(sourcePile.BaseCards) == 0 {
				continue
			}
			fmt.Printf("Can move %#v to pile %#v\n", sourcePile.StackCards[0], targetPile.StackCards)
			possibleMove := Move{
				SourceCard: sourcePile.StackCards[0],
				SourcePileId: sourcePileId,
				TargetPileId: targetPileId,
			}
			if len(targetPile.StackCards) > 0 {
				possibleMove.TargetCard = targetPile.StackCards[len(targetPile.StackCards) - 1]
			}
			possibleMoves = append(possibleMoves, possibleMove)
		}
	}
	if g.Deck.Position == 0 && len(g.Deck.Cards) > 0 {
		g.Deck.Position = 1
		g.Moves++
	}
	for targetPileId := 0; targetPileId < 7; targetPileId++ {
		targetPile := g.Piles[targetPileId]
		currentCard, err := g.Deck.GetCurrentCard()
		if err != nil {
			continue
		}
		if targetPile.CanMoveCardToPile(currentCard) {
			possibleMoves = append(possibleMoves, Move{
				SourceCard: currentCard,
				SourcePileId: 7,
				TargetPileId: targetPileId,
			})
		}
	}
	return possibleMoves
}

func (g *Game) FindAndMakePossibleMoves() bool {
	possibleMoves := g.FindPossibleMoves()

	for _, move := range possibleMoves {
		g.MakeMove(move)
	}

	return len(possibleMoves) > 0
}

func (g *Game) MakeMove(m Move) {
	targetPile := g.Piles[m.TargetPileId]
	if m.SourcePileId == 7 {
		currentCard, err := g.Deck.GetCurrentCard()
		if err != nil {
			return
		}
		if currentCard != m.SourceCard || ! targetPile.CanMoveCardToPile(m.SourceCard) {
			fmt.Print("Cannot make move.\n")
			return
		}
		g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, currentCard)
		g.Deck.PlayCurrentCard()
		return
	}
	sourcePile := g.Piles[m.SourcePileId]
	if len(sourcePile.StackCards) == 0 || !targetPile.CanMoveCardToPile(sourcePile.StackCards[0]) {
		fmt.Print("Cannot make move.\n")
		return
	}
	g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, g.Piles[m.SourcePileId].StackCards...)
	g.Piles[m.SourcePileId].StackCards = g.Piles[m.SourcePileId].StackCards[:0]
	g.Moves++
}
