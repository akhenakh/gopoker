package poker

import "testing"

func TestDeck(t *testing.T) {
	d := NewDeck()
	Equal(t, 52, len(d.cards))
}
