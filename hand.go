package poker

import "fmt"

// Hand is a 5 Cards poker hand
type Hand struct {
	Cards [5]Card
}

// String display a Card as string
func (h *Hand) String() string {
	return fmt.Sprintf("%s %s %s %s %s", h.Cards[0], h.Cards[1], h.Cards[2], h.Cards[3], h.Cards[4])
}
