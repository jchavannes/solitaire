package sol

type Pile struct {
	BaseCards  []Card
	StackCards []Card
}

func (p *Pile) Flip() bool {
	if len(p.StackCards) == 0 && len(p.BaseCards) >= 1 {
		p.StackCards = append(p.StackCards, p.BaseCards[0])
		p.BaseCards = p.BaseCards[1:]
		return true
	}
	return false
}

func (p *Pile) CanMoveCardToPile(card Card) bool {
	pileSize := len(p.StackCards)
	if pileSize == 0 && card.Number == 13 {
		return true
	}
	if pileSize == 0 || card.Number < 2 {
		return false
	}
	stackCard := p.StackCards[pileSize - 1]
	if stackCard.Number == card.Number + 1 && IsOppositeSuits(stackCard.Suit, card.Suit) {
		return true
	}
	return false
}
