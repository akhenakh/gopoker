package poker

import "testing"

func TestCards(t *testing.T) {
	c := &Card{}
	if c.String() != "UndefinedUndefined" {
		t.Fatal("expected undefined cards")
	}

	// -1 isnt possible
	c, err := NewCard(-1, 1)
	if err == nil {
		t.Fatal("expected an error")
	}

	// 0 is undefined
	c, err = NewCard(2, 0)
	if err == nil {
		t.Fatal("expected an error")
	}

	// As is 14
	c, err = NewCard(1, 1)
	if err == nil {
		t.Fatal("expected an error")
	}

	c, err = NewCard(2, 1)
	if err != nil {
		t.Fatal("expected a valid Card")
	}

	if c.String() != "2♦" {
		t.Fatal("expected 2♦ got", c.String())
	}
}

func TestCardsTestFuncs(t *testing.T) {
	c := newCardString("")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCardString("q")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCardString("qa")
	if c != nil {
		t.Fatal("expected a nil card")
	}

	c = newCardString("2♥")
	if c == nil {
		t.Fatal("expected a valid card got nil")
	}

	if c.value != 2 || c.suit != 3 {
		t.Fatal("expected 2♥ got", c)
	}

	c = newCardString("10♣")
	if c.value != 10 || c.suit != 4 {
		t.Fatal("expected 10, 1 got", c.value, c.suit)
	}

	c = newCardString("K♠")
	if c.value != 13 || c.suit != 2 {
		t.Fatal("expected 13, 2 got", c.value, c.suit)
	}
}
