package poker

import "testing"

func TestCards(t *testing.T) {
	c := &Card{}
	if c.String() != "UndefinedUndefined" {
		t.Error("expected undefined cards")
	}

	// -1 isnt possible
	err, c := NewCard(-1, 1)
	if err == nil {
		t.Error("expected an error")
	}

	// 0 is undefined
	err, c = NewCard(2, 0)
	if err == nil {
		t.Error("expected an error")
	}

	// As is 14
	err, c = NewCard(1, 1)
	if err == nil {
		t.Error("expected an error")
	}

	err, c = NewCard(2, 1)
	if err != nil {
		t.Error("expected a valid Card")
	}

	if c.String() != "2♦" {
		t.Error("expected 2♦ got", c.String())
	}
}
