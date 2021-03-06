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
	NoPileCards   []NoPileCard
}

const (
	PileDeck = 7
	PileFoundation = 8
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
		possibleMoves = append(possibleMoves, g.findDeckToFoundationMoves()...)

		possibleMoves = append(possibleMoves, g.findPileToFoundationMoves()...)
		if len(possibleMoves) == 0 {
			possibleMoves = append(possibleMoves, g.findDeckToPileMoves()...)
			if len(possibleMoves) == 0 {
				if g.RecentSubCard {
					g.RecentSubCard = false
				} else {
					subCardMoves := g.findPileToPileMovesWithSubCards()
					possibleMoves = append(possibleMoves, subCardMoves...)
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
		PILES:
		for targetPileId, targetPile := range g.Piles {
			if targetPileId == sourcePileId || ! targetPile.CanMoveCardToPile(sourcePile.StackCards[0]) {
				continue
			}
			if sourcePile.StackCards[0].Number == 13 && len(sourcePile.BaseCards) == 0 {
				continue
			}
			for i, noPileCard := range g.NoPileCards {
				if sourcePile.StackCards[0] == noPileCard.Card && noPileCard.Times > 0 {
					g.NoPileCards[i].Times--
					continue PILES
				}
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
					SourcePileIndex: i,
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
	for i, noPileCard := range g.NoPileCards {
		//fmt.Printf("noPileCard: %d, g.Moves: %d\n", noPileCard, g.Moves)
		//fmt.Printf("g.Deck.Cards[noPileCard]: %#v\n", g.Deck.Cards[noPileCard])
		if currentCard == noPileCard.Card && noPileCard.Times > 0 {
			g.NoPileCards[i].Times--
			return possibleMoves
		}
	}
	if err == nil {
		for targetPileId, targetPile := range g.Piles {
			if targetPileId != PileDeck && targetPile.CanMoveCardToPile(currentCard) {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: currentCard,
					SourcePileId: PileDeck,
					TargetPileId: targetPileId,
				})
				break
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
		index := len(sourcePile.StackCards) - 1
		sourceCard := sourcePile.StackCards[index]
		for _, foundation := range g.Foundations {
			if len(foundation.Cards) == 0 && sourceCard.Number == 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: sourceCard,
					SourcePileId: sourcePileId,
					TargetPileId: PileFoundation,
					SourcePileIndex: index,
				})
				break
			}
			if sourceCard.Suit == foundation.Suit && foundation.Cards[len(foundation.Cards) - 1].Number == sourceCard.Number - 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: sourceCard,
					SourcePileId: sourcePileId,
					TargetPileId: PileFoundation,
					SourcePileIndex: index,
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
					SourcePileId: PileDeck,
					TargetPileId: PileFoundation,
				})
				break
			}
			if len(foundation.Cards) == 0 || foundation.Suit != currentCard.Suit {
				continue
			}
			if foundation.Cards[len(foundation.Cards) - 1].Number == currentCard.Number - 1 {
				possibleMoves = append(possibleMoves, Move{
					SourceCard: currentCard,
					SourcePileId: PileDeck,
					TargetPileId: PileFoundation,
				})
			}
		}
	}
	return possibleMoves
}

func (g *Game) MakeMove(m Move) bool {
	if m.TargetPileId == PileFoundation {
		if m.SourcePileId == PileDeck {
			return g.moveDeckToFoundation(m)
		} else {
			return g.movePileToFoundation(m)
		}
	} else if m.SourcePileId == PileDeck {
		return g.moveDeckToPile(m)
	} else {
		return g.movePileToPile(m)
	}
}

func (g *Game) movePileToPile(m Move) bool {
	targetPile := g.Piles[m.TargetPileId]
	sourcePile := g.Piles[m.SourcePileId]
	emptyPiles := false
	for _, pile := range g.Piles {
		if len(pile.BaseCards) == 0 && len(pile.StackCards) == 0 {
			emptyPiles = true
		}
	}
	if len(sourcePile.BaseCards) == 0 && m.SourcePileIndex == 0 && emptyPiles {
		return false
	}
	sourceIndex := -1
	for i, sourcePileCard := range g.Piles[m.SourcePileId].StackCards {
		if m.SourceCard == sourcePileCard {
			sourceIndex = i
		}
	}
	if sourceIndex == -1 {
		return false
	}
	if len(sourcePile.StackCards) == 0 || ! targetPile.CanMoveCardToPile(m.SourceCard) {
		return false
	}
	g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, g.Piles[m.SourcePileId].StackCards[sourceIndex:]...)
	g.Piles[m.SourcePileId].StackCards = g.Piles[m.SourcePileId].StackCards[:sourceIndex]
	g.Moves++
	return true
}

func (g *Game) moveDeckToPile(m Move) bool {
	targetPile := g.Piles[m.TargetPileId]
	currentCard, err := g.Deck.GetCurrentCard()
	if err != nil || currentCard != m.SourceCard || ! targetPile.CanMoveCardToPile(m.SourceCard) {
		return false
	}
	g.Piles[m.TargetPileId].StackCards = append(g.Piles[m.TargetPileId].StackCards, currentCard)
	g.Deck.PlayCurrentCard()
	g.Moves++
	return true
}

func (g *Game) movePileToFoundation(m Move) bool {
	sourcePile := g.Piles[m.SourcePileId]
	if len(sourcePile.StackCards) == 0 {
		return false
	}
	currentCard := sourcePile.StackCards[len(sourcePile.StackCards) - 1]
	if currentCard != m.SourceCard {
		return false
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
		return true
	}
	return false
}

func (g *Game) moveDeckToFoundation(m Move) bool {
	currentCard, err := g.Deck.GetCurrentCard()
	if err != nil || currentCard != m.SourceCard {
		return false
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
		return true
	}
	return false
}
