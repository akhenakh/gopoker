package poker

import "fmt"

// Hand is a 5 Cards poker hand
type Hand struct {
	Cards [5]Card
	value [6]int8
}

// NewHand is returning a *Hand and evaluate it
func NewHand(Cards [5]Card) (*Hand, error) {
	for i, c := range Cards {
		if !c.valid() {
			return nil, fmt.Errorf("Invalid card at pos %d", i)
		}
	}

	h := &Hand{Cards: Cards}
	h.evaluate()
	return h, nil
}

// evaluate is evaluating the poker hand and store results in value
func (h *Hand) evaluate() {

}

// String display a Card as string
func (h *Hand) String() string {
	return fmt.Sprintf("%s %s %s %s %s", h.Cards[0], h.Cards[1], h.Cards[2], h.Cards[3], h.Cards[4])
}
