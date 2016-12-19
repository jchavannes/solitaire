package sol

import (
	"fmt"
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

func (c *Card) GetString() string {
	return fmt.Sprintf("%2d%c", c.Number, c.Suit[0])
}
