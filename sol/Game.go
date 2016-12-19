package sol

import (
	"fmt"
)

type Game struct {
	Piles        [7][]Card // Original piles
	PileStacks   [7][]Card // Stacks of cards on top of piles
	Foundations  [4][]Card

	Deck         []Card
	DeckPosition int
}

func (g *Game) FlipPiles() {
	for pile := 0; pile < 7; pile++ {
		if len(g.PileStacks[pile]) == 0 && len(g.Piles[pile]) >= 1 {
			g.PileStacks[pile] = append(g.PileStacks[pile], g.Piles[pile][0])
			g.Piles[pile] = g.Piles[pile][1:]
		}
	}
}

func (g *Game) CanMoveCardToStack(card Card, pile int) bool {
	pileSize := len(g.PileStacks[pile])
	if pileSize == 0 && card.Number == 13 {
		return true
	}
	if pileSize == 0 || card.Number < 2 {
		return false
	}
	stackCard := g.PileStacks[pile][pileSize - 1]
	if stackCard.Number == card.Number + 1 && IsOppositeSuits(stackCard.Suit, card.Suit) {
		return true
	}
	return false
}

func (g *Game) OutputMissingCards() {
	for _, suit := range []Suit{Hearts, Diamonds, Spades, Clubs} {
		for number := 1; number <= 13; number++ {
			found := 0
			for pile := 0; pile < 7; pile++ {
				if IsCardInPile(g.Piles[pile], Card{Number: number, Suit: suit}) {
					found++
				}
			}
			if IsCardInPile(g.Deck, Card{Number: number, Suit: suit}) {
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
		if len(g.Piles[pile]) == 0 {
			printString = printString + "  -  "
		} else {
			printString = printString + fmt.Sprintf("  %1d  ", len(g.Piles[pile]))
		}
	}
	println(printString)
	for cardNum := 0; found; cardNum++ {
		found = false
		printString = ""
		for pile := 0; pile < 7; pile++ {
			if len(g.PileStacks[pile]) > cardNum {
				card := g.PileStacks[pile][cardNum]
				printString = printString + fmt.Sprintf(" %2d%c ", card.Number, card.Suit[0])
				found = true
			} else {
				printString = printString + "     "
			}
		}
		println(printString)
	}
}

func (g *Game) FindPossibleMoves() []Move {
	var possibleMoves []Move
	for sourcePileId := 0; sourcePileId < 7; sourcePileId++ {
		sourcePile := g.PileStacks[sourcePileId]
		for targetPileId := 0; targetPileId < 7; targetPileId++ {
			targetPile := g.PileStacks[targetPileId]
			if targetPileId == sourcePileId || len(sourcePile) == 0 || !g.CanMoveCardToStack(sourcePile[0], targetPileId) {
				continue
			}
			if sourcePile[0].Number == 13 && len(g.Piles[sourcePileId]) == 0 {
				continue
			}
			fmt.Printf("Can move %#v to pile %#v\n", sourcePile[0], targetPile)
			possibleMove := Move{
				SourceCard: sourcePile[0],
				SourcePileId: sourcePileId,
				TargetPileId: targetPileId,
			}
			if len(targetPile) > 0 {
				possibleMove.TargetCard = targetPile[len(targetPile) - 1]
			}
			possibleMoves = append(possibleMoves, possibleMove)
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
	targetPile := g.PileStacks[m.TargetPileId]
	sourcePile := g.PileStacks[m.SourcePileId]
	if len(sourcePile) ==  0 {
		fmt.Print("Cannot make move.\n")
		return
	}
	if len(targetPile) > 0 && m.TargetCard != targetPile[len(targetPile) - 1] {
		fmt.Printf("Cannot make move, target card missing: %#v\n", m)
		return
	}
	if len(targetPile) > 0 && m.SourceCard != sourcePile[0] {
		fmt.Printf("Cannot make move, source card missing: %#v\n", m)
		return
	}
	g.PileStacks[m.TargetPileId] = append(g.PileStacks[m.TargetPileId], g.PileStacks[m.SourcePileId]...)
	g.PileStacks[m.SourcePileId] = g.PileStacks[m.SourcePileId][:0]
}
