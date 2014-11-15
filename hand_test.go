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

	// High card
	h, err = newHS("5♠", "3♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != HighCard {
		t.Fatal("expected HandValue to be a High card got value", h.HandValue)
	}

	// Pair
	h, err = newHS("A♠", "A♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != Pair {
		t.Fatal("expected HandValue to be a Pair got value", h.HandValue)
	}
	if h.EvalString() != "Pair of A's" {
		t.Fatal("expected Pair of A's got", h.EvalString())
	}

	h, err = newHS("K♠", "K♥", "J♥", "7♥", "7♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != TwoPairs {
		t.Fatal("expected HandValue to be Two pairs got value", h.HandValue)
	}
	if h.EvalString() != "Two pair K 7" {
		t.Fatal("expected Two pair K 7 got", h.EvalString())
	}

	// Two pairs
	h, err = newHS("7♠", "7♥", "J♥", "A♥", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != TwoPairs {
		t.Fatal("expected HandValue to be Two pairs got value", h.HandValue)
	}
	if h.EvalString() != "Two pair A 7" {
		t.Fatal("expected Two pair A 7 got", h.EvalString())
	}
	h, err = newHS("Q♠", "A♥", "J♥", "A♣", "Q♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != TwoPairs {
		t.Fatal("expected HandValue to be Two pairs got value", h.HandValue)
	}
	if h.EvalString() != "Two pair A Q" {
		t.Fatal("expected Two pair A Q got", h.EvalString())
	}

	// Three
	h, err = newHS("A♠", "A♥", "J♥", "7♣", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != ThreeOfAKind {
		t.Fatal("expected HandValue to be Three of a Kind got value", h.HandValue)
	}
	if h.EvalString() != "Three of a kind A's" {
		t.Fatal("expected Three A high got", h.EvalString())
	}

	// Straight
	h, err = newHS("A♣", "2♣", "3♦", "4♦", "5♥")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != Straight {
		t.Fatal("expected HandValue to be a Straight got value", h.HandValue)
	}
	if h.EvalString() != "Straight 5 high" {
		t.Fatal("expected Straight 5 high got", h.EvalString())
	}

	// Full House
	h, err = newHS("A♦", "A♥", "J♥", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FullHouse {
		t.Fatal("expected HandValue to be Full House got value", h.HandValue)
	}
	if h.EvalString() != "Full house J over A" {
		t.Fatal("expected Full house J over A got", h.EvalString())
	}
	h, err = newHS("7♦", "7♥", "7♠", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FullHouse {
		t.Fatal("expected HandValue to be Full House got value", h.HandValue)
	}
	if h.EvalString() != "Full house 7 over J" {
		t.Fatal("expected Full house 7 over J got", h.EvalString())
	}

	// Flush
	h, err = newHS("7♥", "9♥", "3♥", "4♥", "2♥")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != Flush {
		t.Fatal("expected HandValue to b Flush got value", h.HandValue)
	}
	if h.EvalString() != "9 high flush" {
		t.Fatal("expected 9 high flush got", h.EvalString())
	}

	// Four
	h, err = newHS("A♦", "A♥", "J♥", "A♠", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != FourOfAKind {
		t.Fatal("expected HandValue to be Four of a Kind got value", h.HandValue)
	}
	if h.EvalString() != "Four of a kind A's" {
		t.Fatal("expected Four A high got", h.EvalString())
	}

	// Straight flush
	h, err = newHS("A♦", "2♦", "3♦", "4♦", "5♦")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != StraightFlush {
		t.Fatal("expected HandValue to be a Straight Flush got value", h.HandValue)
	}
	if h.EvalString() != "Straight flush 5 high" {
		t.Fatal("expected Straight flush 5 high got", h.EvalString())
	}

	// straight flush A high
	h, err = newHS("10♦", "J♦", "K♦", "A♦", "Q♦")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != StraightFlush {
		t.Fatal("expected HandValue to be a Straight Flush got value", h.HandValue)
	}
	if h.EvalString() != "Straight flush A high" {
		t.Fatal("expected Straight flush A high got", h.EvalString())
	}

	//High card
	h, err = newHS("10♦", "2♦", "K♦", "5♦", "Q♥")
	if err != nil {
		t.Fatal(err)
	}
	if h.HandValue != HighCard {
		t.Fatal("expected HandValue to be High card got value", h.HandValue)
	}
	if h.EvalString() != "High card" {
		t.Fatal("expected high card got", h.EvalString())
	}
}
