package sol

import (
	"fmt"
	"strings"
	"math/rand"
)

type Game struct {
	Piles         [7]Pile
	Foundations   [4]Foundation
	Deck          Deck
	Moves         int
	SkipPercent   int
	RecentSubCard bool
}

const (
	SourcePile_Deck = 7
	TargetFile_Foundation = 7
)

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
	println("--------")

	printString := ""
	for cardNum := 0; found; cardNum++ {
		found = false
		printString = ""
		for _, foundation := range g.Foundations {
			if len(foundation.Cards) > cardNum {
				card := foundation.Cards[cardNum]
				printString = printString + fmt.Sprintf(" %3s ", card.GetString())
				found = true
			} else {
				printString = printString + "     "
			}
		}
		println(printString)
	}

	printString = ""
	for _, pile := range g.Piles {
		if len(pile.BaseCards) == 0 {
			printString = printString + "  -  "
		} else {
			printString = printString + fmt.Sprintf("  %1d  ", len(pile.BaseCards))
		}
	}
	println(printString)

	found = true
	for cardNum := 0; found; cardNum++ {
		found = false
		printString = ""
		for _, pile := range g.Piles {
			if len(pile.StackCards) > cardNum {
				card := pile.StackCards[cardNum]
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
		if len(g.Deck.Cards) > 0 {
			g.Deck.NextCard()
			g.Moves++
			currentCard, err = g.Deck.GetCurrentCard()
		} else {
			println("Deck empty")
		}
	}
	if err == nil {
		fmt.Printf("Deck: %d of %d ( %s )\n", g.Deck.Position, len(g.Deck.Cards), currentCard.GetString())
		printString = ""
		var cardStrings []string
		for _, card := range g.Deck.Cards {
			cardStrings = append(cardStrings, card.GetString())
		}
		println(strings.Join(cardStrings, ", "))
	}
	fmt.Printf("Moves: %d\n", g.Moves)
}

func (g *Game) OutputGameSimple() {
	fmt.Printf("Win?: %t\n", g.IsGameCompleted())
	fmt.Printf("Moves: %d\n", g.Moves)
}

func (g *Game) IsGameCompleted() bool {
	for _, foundation := range g.Foundations {
		if len(foundation.Cards) < 13 {
			return false
		}
	}
	return true
}

func (g *Game) SetSkipPercent(skipPercent int) {
	g.SkipPercent = skipPercent
}

/*
- 4 Types of moves:

  1) Pile -> Pile
  2) Deck -> Pile
  3) Pile -> Foundation
  4) Deck -> Foundation

- Moving to a pile either stacks on top of current card or creates a new stack.

- Moving to a foundation either stacks on top of current card or creates a new foundation.
 */
func (g *Game) FindAndMakePossibleMoves() bool {
	possibleMoves := g.FindPossibleMoves()

	for _, move := range possibleMoves {
		g.MakeMove(move)
	}

	return len(possibleMoves) > 0
}

func (g *Game) FindPossibleMoves() []Move {
	if g.Deck.Position == 0 && len(g.Deck.Cards) > 0 {
		g.Deck.Position = 1
		g.Moves++
	}

	var possibleMoves []Move

	possibleMoves = append(possibleMoves, g.findPileToPileMoves()...)
	if len(possibleMoves) == 0 {
		possibleMoves = append(possibleMoves, g.findDeckToPileMoves()...)

		possibleMoves = append(possibleMoves, g.findPileToFoundationMoves()...)
		if len(possibleMoves) == 0 {
			possibleMoves = append(possibleMoves, g.findDeckToFoundationMoves()...)
			if len(possibleMoves) == 0 && rand.Intn(100) < 20 {
				if g.RecentSubCard {
					g.RecentSubCard = false
				} else {
					possibleMoves = append(possibleMoves, g.findPileToPileMovesWithSubCards()...)
					g.RecentSubCard = true
				}
			}
		}
	}

	for i := range possibleMoves {
		if g.SkipPercent > rand.Intn(100) {
			if (i >= len(possibleMoves)) {
				possibleMoves = possibleMoves[:i]
			} else {
				possibleMoves = append(possibleMoves[:i], possibleMoves[i + 1:]...)
			}
		}
	}

	return possibleMoves
}

func (g *Game) findPileToPileMoves() []Move {
	var possibleMoves []Move
	for sourcePileId, sourcePile := range g.Piles {
		if len(sourcePile.StackCards) == 0 {
			continue
		}
		emptyPiles := false
		for _, pile := range g.Piles {
			if len(pile.BaseCards) == 0 && len(pile.StackCards) == 0 {
				emptyPiles = true
			}
		}
		if len(sourcePile.BaseCards) == 0 && emptyPiles {
			continue
		}
		for targetPileId, targetPile := range g.Piles {
			if targetPileId == sourcePileId || ! targetPile.CanMoveCardToPile(sourcePile.StackCards[0]) {
				continue
			}
			if sourcePile.StackCards[0].Number == 13 && len(sourcePile.BaseCards) == 0 {
				continue
			}
			//fmt.Printf("Can move %#v to pile %#v\n", sourcePile.StackCards[0], targetPile.StackCards)
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
	return possibleMoves
}

func (g *Game) findPileToPileMovesWithSubCards() []Move {
	var possibleMoves []Move
	for sourcePileId, sourcePile := range g.Piles {
		if len(sourcePile.StackCards) < 2 {
			continue
		}
		for targetPileId, targetPile := range g.Piles {
			if targetPileId == sourcePileId {
				continue
			}
			for i, currentCard := range sourcePile.StackCards {
				if i == 0 || ! targetPile.CanMoveCardToPile(currentCard) {
					continue
				}
				possibleMove := Move{
					SourceCard: currentCard,
					SourcePileId: sourcePileId,
					TargetPileId: targetPileId,
				}
				if len(targetPile.StackCards) > 0 {
					possibleMove.TargetCard = targetPile.StackCards[len(targetPile.StackCards) - 1]
				}
				possibleMoves = append(possibleMoves, possibleMove)
			}
		}
	}
	return possibleMoves
}

func (g *Game) findDeckToPileMoves() []Move {
	var possibleMoves []Move
	currentCard, err := g.Deck.GetCurrentCard()
	if err == nil {
		for targetPileId, targetPile := range g.Piles {
			if targetPile.CanMoveCardToPile(currentCard) {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: currentCard,
					SourcePileId: SourcePile_Deck,
					TargetPileId: targetPileId,
				})
			}
		}
	}
	return possibleMoves
}

func (g *Game) findPileToFoundationMoves() []Move {
	var possibleMoves []Move
	for sourcePileId, sourcePile := range g.Piles {
		if len(sourcePile.StackCards) == 0 {
			continue
		}
		sourceCard := sourcePile.StackCards[len(sourcePile.StackCards) - 1]
		for _, foundation := range g.Foundations {
			if len(foundation.Cards) == 0 && sourceCard.Number == 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: sourceCard,
					SourcePileId: sourcePileId,
					TargetPileId: TargetFile_Foundation,
				})
				break
			}
			if sourceCard.Suit == foundation.Suit && foundation.Cards[len(foundation.Cards) - 1].Number == sourceCard.Number - 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: sourceCard,
					SourcePileId: sourcePileId,
					TargetPileId: TargetFile_Foundation,
				})
				break
			}
		}
	}
	return possibleMoves
}

func (g *Game) findDeckToFoundationMoves() []Move {
	var possibleMoves []Move
	currentCard, err := g.Deck.GetCurrentCard()
	if err == nil {
		for _, foundation := range g.Foundations {
			if len(foundation.Cards) == 0 && currentCard.Number == 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: currentCard,
					SourcePileId: SourcePile_Deck,
					TargetPileId: TargetFile_Foundation,
				})
				break
			}
			if len(foundation.Cards) == 0 || foundation.Suit != currentCard.Suit {
				continue
			}
			if foundation.Cards[len(foundation.Cards) - 1].Number == currentCard.Number - 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: currentCard,
					SourcePileId: SourcePile_Deck,
					TargetPileId: TargetFile_Foundation,
				})
			}
		}
	}
	return possibleMoves
}

func (g *Game) MakeMove(m Move) {
	if m.TargetPileId == TargetFile_Foundation {
		if m.SourcePileId == SourcePile_Deck {
			g.moveDeckToFoundation(m)
		} else {
			g.movePileToFoundation(m)
		}
	} else if m.SourcePileId == SourcePile_Deck {
		g.moveDeckToPile(m)
		return
	} else {
		g.movePileToPile(m)
	}
}

func (g *Game) movePileToPile(m Move) {
	targetPile := g.Piles[m.TargetPileId]
	sourcePile := g.Piles[m.SourcePileId]
	emptyPiles := false
	for _, pile := range g.Piles {
		if len(pile.BaseCards) == 0 && len(pile.StackCards) == 0 {
			emptyPiles = true
		}
	}
	if len(sourcePile.BaseCards) == 0 && emptyPiles {
		return
	}
	sourceIndex := -1
	for i, sourcePileCard := range g.Piles[m.SourcePileId].StackCards {
		if m.SourceCard == sourcePileCard {
			sourceIndex = i
		}
	}
	if sourceIndex == -1 {
		return
	}
	if len(sourcePile.StackCards) == 0 || ! targetPile.CanMoveCardToPile(m.SourceCard) {
		return
	}
	g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, g.Piles[m.SourcePileId].StackCards[sourceIndex:]...)
	g.Piles[m.SourcePileId].StackCards = g.Piles[m.SourcePileId].StackCards[:sourceIndex]
	g.Moves++
}

func (g *Game) moveDeckToPile(m Move) {
	targetPile := g.Piles[m.TargetPileId]
	currentCard, err := g.Deck.GetCurrentCard()
	if err != nil || currentCard != m.SourceCard || ! targetPile.CanMoveCardToPile(m.SourceCard) {
		return
	}
	g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, currentCard)
	g.Deck.PlayCurrentCard()
	g.Moves++
}

func (g *Game) movePileToFoundation(m Move) {
	sourcePile := g.Piles[m.SourcePileId]
	if len(sourcePile.StackCards) == 0 {
		return
	}
	currentCard := sourcePile.StackCards[len(sourcePile.StackCards) - 1]
	if currentCard != m.SourceCard {
		return
	}
	for foundationId, foundation := range g.Foundations {
		if currentCard.Number == 1 && foundation.Suit == "" {
			g.Foundations[foundationId].Suit = currentCard.Suit
		} else if foundation.Suit != currentCard.Suit || foundation.Cards[len(foundation.Cards) - 1].Number != currentCard.Number - 1 {
			continue
		}
		g.Foundations[foundationId].Cards = append(g.Foundations[foundationId].Cards, currentCard)
		if currentCard.Number == 8 && currentCard.Suit == Hearts {
			//fmt.Printf("Stack: %#v\n", g.Piles[m.SourcePileId].StackCards)
		}
		g.Piles[m.SourcePileId].StackCards = sourcePile.StackCards[:len(sourcePile.StackCards) - 1]
		if currentCard.Number == 8 && currentCard.Suit == Hearts {
			//fmt.Printf("Stack: %#v\n", g.Piles[m.SourcePileId].StackCards)
		}
		g.Moves++
		break
	}
}

func (g *Game) moveDeckToFoundation(m Move) {
	currentCard, err := g.Deck.GetCurrentCard()
	if err != nil || currentCard != m.SourceCard {
		return
	}
	for foundationId, foundation := range g.Foundations {
		if currentCard.Number == 1 && foundation.Suit == "" {
			g.Foundations[foundationId].Suit = currentCard.Suit
		} else if foundation.Suit != currentCard.Suit || foundation.Cards[len(foundation.Cards) - 1].Number != currentCard.Number - 1 {
			continue
		}
		g.Foundations[foundationId].Cards = append(g.Foundations[foundationId].Cards, currentCard)
		g.Deck.PlayCurrentCard()
		g.Moves++
		break
	}
}
