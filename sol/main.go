package sol

func IsOppositeSuits(s1 Suit, s2 Suit) bool {
	if (s1 == Hearts || s1 == Diamonds) && (s2 == Spades || s2 == Clubs) {
		return true
	}
	if (s1 == Spades || s1 == Clubs) && (s2 == Hearts || s2 == Diamonds) {
		return true
	}
	return false
}

func IsCardInPile(stack []Card, match Card) bool {
	for _, card := range stack {
		if card.Number == match.Number && card.Suit == match.Suit {
			return true
		}
	}
	return false
}

func GetGame1() *Game {
	return &Game{
		Piles: [7]Pile{
			{BaseCards: []Card{
				{Number: 11, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 2, Suit: Diamonds},
				{Number: 12, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 11, Suit: Hearts},
				{Number: 4, Suit: Spades},
				{Number: 12, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 3, Suit: Spades},
				{Number: 13, Suit: Diamonds},
				{Number: 6, Suit: Spades},
				{Number: 1, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 6, Suit: Diamonds},
				{Number: 5, Suit: Hearts},
				{Number: 2, Suit: Clubs},
				{Number: 1, Suit: Hearts},
				{Number: 5, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 7, Suit: Spades},
				{Number: 3, Suit: Hearts},
				{Number: 4, Suit: Hearts},
				{Number: 7, Suit: Diamonds},
				{Number: 3, Suit: Clubs},
				{Number: 8, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 13, Suit: Hearts},
				{Number: 5, Suit: Clubs},
				{Number: 13, Suit: Spades},
				{Number: 10, Suit: Spades},
				{Number: 9, Suit: Clubs},
				{Number: 3, Suit: Diamonds},
				{Number: 11, Suit: Spades},
			}},
		},
		Deck: Deck{Cards: []Card{
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
		}},
	}
}

func GetGame2() *Game {
	return &Game{
		Piles: [7]Pile{
			{BaseCards: []Card{
				{Number: 1, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 1, Suit: Spades},
				{Number: 10, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 12, Suit: Clubs},
				{Number: 5, Suit: Spades},
				{Number: 6, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 11, Suit: Spades},
				{Number: 7, Suit: Spades},
				{Number: 3, Suit: Diamonds},
				{Number: 11, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 9, Suit: Spades},
				{Number: 9, Suit: Clubs},
				{Number: 6, Suit: Clubs},
				{Number: 7, Suit: Clubs},
				{Number: 13, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 8, Suit: Spades},
				{Number: 8, Suit: Clubs},
				{Number: 8, Suit: Diamonds},
				{Number: 10, Suit: Clubs},
				{Number: 4, Suit: Diamonds},
				{Number: 6, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 8, Suit: Hearts},
				{Number: 13, Suit: Clubs},
				{Number: 3, Suit: Hearts},
				{Number: 5, Suit: Hearts},
				{Number: 1, Suit: Clubs},
				{Number: 7, Suit: Hearts},
				{Number: 11, Suit: Clubs},
			}},
		},
		Deck: Deck{Cards: []Card{
			{Number: 4, Suit: Spades},
			{Number: 4, Suit: Hearts},
			{Number: 3, Suit: Clubs},
			{Number: 10, Suit: Diamonds},
			{Number: 9, Suit: Diamonds},
			{Number: 4, Suit: Clubs},
			{Number: 9, Suit: Hearts},
			{Number: 5, Suit: Diamonds},
			{Number: 6, Suit: Hearts},
			{Number: 11, Suit: Hearts},
			{Number: 2, Suit: Diamonds},
			{Number: 13, Suit: Hearts},
			{Number: 1, Suit: Hearts},
			{Number: 2, Suit: Hearts},
			{Number: 5, Suit: Clubs},
			{Number: 3, Suit: Spades},
			{Number: 10, Suit: Hearts},
			{Number: 12, Suit: Spades},
			{Number: 2, Suit: Clubs},
			{Number: 12, Suit: Diamonds},
			{Number: 2, Suit: Spades},
			{Number: 13, Suit: Spades},
			{Number: 12, Suit: Hearts},
			{Number: 7, Suit: Diamonds},
		}},
	}
}

// Daily challenge 2017-04-08
// High score: 96
// #2 high score: 115
// Lowest high score: 116
// Personal score: 119
func GetGame3() *Game {
	return &Game{
		Piles: [7]Pile{
			{BaseCards: []Card{
				{Number: 8, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 9, Suit: Clubs},
				{Number: 4, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 5, Suit: Diamonds},
				{Number: 3, Suit: Spades},
				{Number: 1, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 11, Suit: Hearts},
				{Number: 9, Suit: Diamonds},
				{Number: 12, Suit: Diamonds},
				{Number: 10, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 3, Suit: Hearts},
				{Number: 10, Suit: Hearts},
				{Number: 7, Suit: Clubs},
				{Number: 7, Suit: Diamonds},
				{Number: 2, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 6, Suit: Diamonds},
				{Number: 2, Suit: Diamonds},
				{Number: 11, Suit: Spades},
				{Number: 5, Suit: Clubs},
				{Number: 7, Suit: Hearts},
				{Number: 3, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 12, Suit: Clubs},
				{Number: 6, Suit: Spades},
				{Number: 12, Suit: Spades},
				{Number: 6, Suit: Hearts},
				{Number: 9, Suit: Spades},
				{Number: 10, Suit: Diamonds},
				{Number: 13, Suit: Diamonds},
			}},
		},
		Deck: Deck{Cards: []Card{
			{Number: 2, Suit: Spades},
			{Number: 8, Suit: Spades},
			{Number: 7, Suit: Spades},
			{Number: 11, Suit: Clubs},
			{Number: 5, Suit: Spades},
			{Number: 12, Suit: Hearts},
			{Number: 4, Suit: Spades},
			{Number: 13, Suit: Spades},
			{Number: 1, Suit: Hearts},
			{Number: 11, Suit: Diamonds},
			{Number: 5, Suit: Hearts},
			{Number: 13, Suit: Hearts},
			{Number: 2, Suit: Hearts},
			{Number: 8, Suit: Diamonds},
			{Number: 10, Suit: Spades},
			{Number: 8, Suit: Hearts},
			{Number: 3, Suit: Diamonds},
			{Number: 6, Suit: Clubs},
			{Number: 4, Suit: Diamonds},
			{Number: 1, Suit: Spades},
			{Number: 4, Suit: Hearts},
			{Number: 1, Suit: Clubs},
			{Number: 13, Suit: Clubs},
			{Number: 9, Suit: Hearts},
		}},
	}
}

// Daily challenge 2017-04-07
// High score: 98
// #2 high score: 108
// Lowest high score: 109
// Personal score: 117
func GetGame4() *Game {
	return &Game{
		Piles: [7]Pile{
			{BaseCards: []Card{
				{Number: 7, Suit: Hearts},
			}},
			{BaseCards: []Card{
				{Number: 2, Suit: Spades},
				{Number: 4, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 8, Suit: Spades},
				{Number: 9, Suit: Clubs},
				{Number: 1, Suit: Clubs},
			}},
			{BaseCards: []Card{
				{Number: 10, Suit: Clubs},
				{Number: 13, Suit: Clubs},
				{Number: 3, Suit: Spades},
				{Number: 11, Suit: Diamonds},
			}},
			{BaseCards: []Card{
				{Number: 2, Suit: Clubs},
				{Number: 13, Suit: Hearts},
				{Number: 7, Suit: Clubs},
				{Number: 1, Suit: Hearts},
				{Number: 13, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 13, Suit: Diamonds},
				{Number: 7, Suit: Diamonds},
				{Number: 6, Suit: Spades},
				{Number: 3, Suit: Clubs},
				{Number: 3, Suit: Hearts},
				{Number: 9, Suit: Spades},
			}},
			{BaseCards: []Card{
				{Number: 5, Suit: Diamonds},
				{Number: 11, Suit: Spades},
				{Number: 4, Suit: Clubs},
				{Number: 5, Suit: Hearts},
				{Number: 6, Suit: Clubs},
				{Number: 7, Suit: Spades},
				{Number: 6, Suit: Hearts},
			}},
		},
		Deck: Deck{Cards: []Card{
			{Number: 10, Suit: Diamonds},
			{Number: 4, Suit: Hearts},
			{Number: 12, Suit: Clubs},
			{Number: 9, Suit: Diamonds},
			{Number: 4, Suit: Spades},
			{Number: 11, Suit: Hearts},
			{Number: 6, Suit: Diamonds},
			{Number: 1, Suit: Spades},
			{Number: 10, Suit: Hearts},
			{Number: 5, Suit: Spades},
			{Number: 12, Suit: Spades},
			{Number: 3, Suit: Diamonds},
			{Number: 5, Suit: Clubs},
			{Number: 8, Suit: Hearts},
			{Number: 12, Suit: Diamonds},
			{Number: 11, Suit: Clubs},
			{Number: 2, Suit: Hearts},
			{Number: 2, Suit: Diamonds},
			{Number: 10, Suit: Spades},
			{Number: 9, Suit: Hearts},
			{Number: 1, Suit: Diamonds},
			{Number: 8, Suit: Clubs},
			{Number: 8, Suit: Diamonds},
			{Number: 12, Suit: Hearts},
		}},
	}
}
