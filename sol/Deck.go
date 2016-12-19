package sol

import "errors"

type Deck struct {
	Cards    []Card
	Position int
}

func (d *Deck) GetCurrentCard() (Card, error) {
	if d.Position == 0 || d.Position > len(d.Cards) {
		return Card{}, errors.New("Unable to find card.")
	}
	return d.Cards[d.Position - 1], nil
}

func (d *Deck) PlayCurrentCard() {
	d.Cards = append(d.Cards[:d.Position - 1], d.Cards[d.Position:]...)
	d.Position--
}

func (d *Deck) NextCard() {
	d.Position++
	if d.Position > len(d.Cards) {
		d.Position = 0
	}
}
