package main

import (
	"fmt"
)

type Suit string

const (
	Hearts Suit = "hearts"
	Diamonds Suit = "diamonds"
	Spades Suit = "spades"
	Clubs Suit = "clubs"
)

type Card struct {
	/**
	 * A = 1
	 * 2 - 10
	 * J = 11
	 * Q = 12
	 * K = 13
	 */
	Number int
	Suit   Suit
}

type SolitaireGame struct {
	Piles        [7][]Card // Original piles
	PileStacks   [7][]Card // Stacks of cards on top of piles
	Foundations  [4][]Card

	Deck         []Card
	DeckPosition int
}

func (g *SolitaireGame) FlipPiles() {
	for pile := 0; pile < 7; pile++ {
		if len(g.PileStacks[pile]) == 0 && len(g.Piles[pile]) >= 1 {
			g.PileStacks[pile] = append(g.PileStacks[pile], g.Piles[pile][0])
			g.Piles[pile] = g.Piles[pile][1:]
		}
	}
}

func (g *SolitaireGame) CanMoveCardToStack(card Card, pile int) bool {
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

func (g *SolitaireGame) OutputMissingCards() {
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

func (g *SolitaireGame) OutputGame() {
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

func (g *SolitaireGame) FindPossibleMoves() []SolitaireMove {
	var possibleMoves []SolitaireMove
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
			possibleMove := SolitaireMove{
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

func (g *SolitaireGame) FindAndMakePossibleMoves() bool {
	possibleMoves := g.FindPossibleMoves()

	for _, move := range possibleMoves {
		g.MakeMove(move)
	}

	return len(possibleMoves) > 0
}

func (g *SolitaireGame) MakeMove(m SolitaireMove) {
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

type SolitaireMove struct {
	SourceCard   Card
	TargetCard   Card
	SourcePileId int
	TargetPileId int
}

func IsOppositeSuits(s1 Suit, s2 Suit) bool {
	if (s1 == Hearts || s1 == Diamonds) && (s2 == Spades || s2 == Clubs) {
		return true
	}
	if (s1 == Spades || s1 == Clubs) && (s2 == Hearts || s2 == Diamonds) {
		return true
	}
	return false
}

func main() {
	println("Solitaire.")
	game := getSampleGame()
	game.OutputMissingCards()

	game.FlipPiles()
	game.OutputGame()

	for i := 0; i < 10; i++ {
		if ! game.FindAndMakePossibleMoves() {
			println("No new moves.")
			break
		}
		game.FlipPiles()
		game.OutputGame()
	}

	//fmt.Printf("%#v\n", game)
}

func IsCardInPile(stack []Card, match Card) bool {
	for _, card := range stack {
		if card.Number == match.Number && card.Suit == match.Suit {
			return true
		}
	}
	return false
}

func getSampleGame() *SolitaireGame {
	return &SolitaireGame{
		Piles: [7][]Card{
			[]Card{
				{Number: 11, Suit: Clubs},
			},
			[]Card{
				{Number: 2, Suit: Diamonds},
				{Number: 12, Suit: Spades},
			},
			[]Card{
				{Number: 11, Suit: Hearts},
				{Number: 4, Suit: Spades},
				{Number: 12, Suit: Diamonds},
			},
			[]Card{
				{Number: 3, Suit: Spades},
				{Number: 13, Suit: Diamonds},
				{Number: 6, Suit: Spades},
				{Number: 1, Suit: Diamonds},
			},
			[]Card{
				{Number: 6, Suit: Diamonds},
				{Number: 5, Suit: Hearts},
				{Number: 2, Suit: Clubs},
				{Number: 1, Suit: Hearts},
				{Number: 5, Suit: Spades},
			},
			[]Card{
				{Number: 7, Suit: Spades},
				{Number: 3, Suit: Hearts},
				{Number: 4, Suit: Hearts},
				{Number: 7, Suit: Diamonds},
				{Number: 3, Suit: Clubs},
				{Number: 8, Suit: Clubs},
			},
			[]Card{
				{Number: 13, Suit: Hearts},
				{Number: 5, Suit: Clubs},
				{Number: 11, Suit: Spades},
				{Number: 10, Suit: Spades},
				{Number: 9, Suit: Clubs},
				{Number: 3, Suit: Diamonds},
				{Number: 13, Suit: Spades},
			},
		},
		Deck: []Card{
			{Number: 4, Suit: Diamonds},
			{Number: 8, Suit: Spades},
			{Number: 7, Suit: Clubs},
			{Number: 12, Suit: Hearts},
			{Number: 8, Suit: Diamonds},
			{Number: 9, Suit: Hearts},
			{Number: 2, Suit: Spades},
			{Number: 4, Suit: Clubs},
			{Number: 10, Suit: Clubs},
			{Number: 12, Suit: Clubs},
			{Number: 5, Suit: Diamonds},
			{Number: 7, Suit: Hearts},
			{Number: 10, Suit: Diamonds},
			{Number: 9, Suit: Diamonds},
			{Number: 9, Suit: Spades},
			{Number: 6, Suit: Hearts},
			{Number: 8, Suit: Hearts},
			{Number: 1, Suit: Clubs},
			{Number: 1, Suit: Spades},
			{Number: 6, Suit: Clubs},
			{Number: 10, Suit: Hearts},
			{Number: 13, Suit: Clubs},
			{Number: 2, Suit: Hearts},
			{Number: 11, Suit: Diamonds},
		},
	}
}

