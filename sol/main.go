package sol

import (
	"strconv"
)

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

func GetGame5() *Game {
	return convertGame("315371647392J3112223K3J2548333T414828452T351J121K2Q241324293449163341343T2Q37261Q481K1Q194J4K424127462T1")
}

func GetGame6() *Game {
	return convertGame("7391847443T2T144711483J39341Q4223352Q261J42432J211312192Q163535113K172Q394K23482621242J1K423T364K35481T4")
}

func GetGame7() *Game {
	return convertGame("Q132K46444528474T2Q2633442K353T43114625154T34373J33391218313Q361J28182K1T1419294J11124237172Q49322K212J4")
}

func convertGame(s string) *Game {
	game := Game{}
	row := 0
	pile := 0
	for i := 0; i < len(s) / 2; i++ {
		card := getCardFromCode(s[i*2:i*2+2])
		if row < 7 {
			game.Piles[pile].BaseCards = append([]Card{card}, game.Piles[pile].BaseCards...)
			pile++
			if pile > 6 {
				row++
				pile = row
			}
		} else {
			game.Deck.Cards = append([]Card{card}, game.Deck.Cards...)
		}
	}
	return &game
}

func getCardFromCode(c string) Card {
	numberCode := c[0:1]
	var number int
	switch numberCode {
	case "T":
		number = 10
		break
	case "J":
		number = 11
		break
	case "Q":
		number = 12
		break
	case "K":
		number = 13
		break
	default:
		number, _ = strconv.Atoi(numberCode)
	}
	suitCode := c[1:2]
	var suit Suit
	switch suitCode {
	case "1":
		suit = Spades
		break
	case "2":
		suit = Diamonds
		break
	case "3":
		suit = Clubs
		break
	case "4":
		suit = Hearts
	}
	card := Card{
		Number: number,
		Suit: suit,
	}
	return card
}
