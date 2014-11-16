package poker

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	cards []Card
}

// NewDeck initializes, shuffles and returns a Deck
func NewDeck() *Deck {
	cards := make([]Card, 52)
	for i := 0; i < 53; i++ {
		c, _ := NewCard(i)
		cards[i] = *c
	}
	d := &Deck{cards: cards}
	return d
}

// Shuffle will shuffles the Deck's cards
func Shuffle() {
	for ii, r := range rand.Perm(52) {
		fmt.Printf("%d is %d\n", ii, r)
	}
}
