package poker

import "testing"

func TestHands(t *testing.T) {
	cards := [5]Card{*newCS("9♥"), *newCS("2♠"), *newCS("A♠"), *newCS("9♣"), *newCS("6♠")}
	h, err := NewHand(cards)
	if err != nil {
		t.Fatal(err)
	}
	if h.String() != "9♥ 2♠ A♠ 9♣ 6♠" {
		t.Fatal("expected 9♥ 2♠ A♠ 9♣ 6♠ got", h.String())
	}

	h, err = newHS("A♠", "2♥", "J♥", "7♣", "A♠")
	if err == nil {
		t.Fatal("expected duplicates")
	}

	h, err = newHS("A♠", "2♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}

	if h.String() != "A♠ 2♥ J♥ 7♣ Q♠" {
		t.Fatal("expected A♠ 2♥ J♥ 7♣ Q♠ got", h.String())
	}

	h, err = newHS("A♠", "A♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != Pair {
		t.Fatal("expected a Pair got value", h.HandValue)
	}

	h, err = newHS("A♠", "A♥", "J♥", "7♥", "7♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != TwoPairs {
		t.Fatal("expected Two pairs got value", h.HandValue)
	}

	h, err = newHS("A♠", "A♥", "J♥", "7♣", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != ThreeOfAKind {
		t.Fatal("expected Three of a Kind got value", h.HandValue)
	}

	h, err = newHS("A♦", "A♥", "J♥", "A♠", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FourOfAKind {
		t.Fatal("expected Four of a Kind got value", h.HandValue)
	}

	h, err = newHS("A♦", "A♥", "J♥", "A♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FullHouse {
		t.Fatal("expected Full House got value", h.HandValue)
	}

	h, err = newHS("A♦", "A♥", "J♥", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FullHouse {
		t.Fatal("expected Full House got value", h.HandValue)
	}

}
