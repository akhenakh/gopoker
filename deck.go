package poker

import (
	"crypto/rand"
	. "math/rand"
)

type Deck struct {
	cards []Card
}

func init() {
	b := make([]byte, 1)
	rand.Read(b)
	Seed(int64(b[0]))
}

// NewDeck initializes, shuffles and returns a Deck
func NewDeck() *Deck {
	cards := make([]Card, 52)
	for i := 1; i < 53; i++ {
		c, _ := NewCard(i)
		cards[i-1] = *c
	}
	d := &Deck{cards: cards}
	d.Shuffle()
	return d
}

// Shuffle will shuffles the Deck's cards
func (d *Deck) Shuffle() {
	for i, j := range Perm(52) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
