package poker

import (
	"testing"
)

func TestCards(t *testing.T) {
	c := &Card{}
	if c.String() != "Undefined" {
		t.Fatal("expected undefined card got", c.String())
	}

	if c.valid() {
		t.Fatal("expected invalid card")
	}

	// -1 isnt possible
	c, err := NewCard(-1)
	if err == nil {
		t.Fatal("expected an error")
	}

	// 0 is undefined
	c, err = NewCard(0)
	if err == nil {
		t.Fatal("expected an error")
	}

	c, err = NewCard(2)
	if err != nil {
		t.Fatal("expected a valid Card")
	}
	Equal(t, "2♦", c.String())

	c, err = NewCard(15)
	if err != nil {
		t.Fatal("expected a valid Card")
	}
	Equal(t, 15, c.value)
	Equal(t, "2♠", c.String())

	// Ace are special case
	c, err = NewCard(1)
	if err != nil {
		t.Fatal("expected a valid Card")
	}
	Equal(t, 14, c.Value())
	Equal(t, 0, c.Suit())
}

func TestCardsTestFuncs(t *testing.T) {
	c := newCS("")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCS("q")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCS("qa")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCS("2♠")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}

	Equal(t, 15, c.value)

	Equal(t, 2, c.Value())

	c = newCS("A♠")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}

	Equal(t, 14, c.Value())

	Equal(t, 1, c.Suit())

	c = newCS("10♣")
	if c.Value() != 10 || c.Suit() != 3 {
		t.Fatal("expected 10, 1 got", c.Value(), c.Suit())
	}

	c = newCS("K♠")
	if c.Value() != 13 || c.Suit() != 1 {
		t.Fatal("expected 13, 1 got", c.Value(), c.Suit())
	}

	c = newCS("A♥")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}

	if c.Value() != 14 {
		t.Fatal("expected 14 got", c.Value())
	}

	c = newCS("K♣")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}
	Equal(t, 52, c.value)

	c = newCS("K♦")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}
	Equal(t, 13, c.Value())
	Equal(t, 0, c.Suit())
}
