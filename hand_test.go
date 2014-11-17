package poker

import (
	"path"
	"reflect"
	"runtime"
	"testing"
)

func Equal(t *testing.T, expected, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("%s:%d expected %v got %v", path.Base(file), line, expected, got)
		t.FailNow()
	}
}

func TestHands(t *testing.T) {
	cards := [5]Card{*newCS("9♥"), *newCS("2♠"), *newCS("A♠"), *newCS("9♣"), *newCS("6♠")}
	h, err := NewHand(cards)
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, "9♥ 2♠ A♠ 9♣ 6♠", h.String())

	h, err = newHS("A♠", "2♥", "J♥", "7♣", "A♠")
	if err == nil {
		t.Fatal("expected duplicates")
	}

	h, err = newHS("A♠", "2♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, "A♠ 2♥ J♥ 7♣ Q♠", h.String())

	// High card
	h, err = newHS("5♠", "3♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, HighCard, h.HandValue)

	// Pair
	h, err = newHS("A♠", "A♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Pair, h.HandValue)
	Equal(t, "Pair of A's", h.EvalString())

	h, err = newHS("K♠", "K♥", "J♥", "7♥", "7♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, TwoPairs, h.HandValue)
	Equal(t, "Two pair K 7", h.EvalString())

	// Two pairs
	h, err = newHS("7♠", "7♥", "J♥", "A♥", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, TwoPairs, h.HandValue)
	Equal(t, "Two pair A 7", h.EvalString())

	h, err = newHS("Q♠", "A♥", "J♥", "A♣", "Q♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, TwoPairs, h.HandValue)
	Equal(t, "Two pair A Q", h.EvalString())

	// Three
	h, err = newHS("A♠", "A♥", "J♥", "7♣", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, ThreeOfAKind, h.HandValue)
	Equal(t, "Three of a kind A's", h.EvalString())

	// Straight
	h, err = newHS("A♣", "2♣", "3♦", "4♦", "5♥")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Straight, h.HandValue)
	Equal(t, "Straight 5 high", h.EvalString())

	// Full House
	h, err = newHS("A♦", "A♥", "J♥", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, FullHouse, h.HandValue)
	Equal(t, "Full house J over A", h.EvalString())

	h, err = newHS("7♦", "7♥", "7♠", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, FullHouse, h.HandValue)
	Equal(t, "Full house 7 over J", h.EvalString())

	// Flush
	h, err = newHS("7♥", "9♥", "3♥", "4♥", "2♥")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Flush, h.HandValue)
	Equal(t, "9 high flush", h.EvalString())

	// Four
	h, err = newHS("A♦", "A♥", "J♥", "A♠", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, FourOfAKind, h.HandValue)
	Equal(t, "Four of a kind A's", h.EvalString())

	// Straight flush
	h, err = newHS("A♦", "2♦", "3♦", "4♦", "5♦")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, StraightFlush, h.HandValue)
	Equal(t, "Straight flush 5 high", h.EvalString())

	// straight flush A high
	h, err = newHS("10♦", "J♦", "K♦", "A♦", "Q♦")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, StraightFlush, h.HandValue)
	Equal(t, "Straight flush A high", h.EvalString())

	//High card
	h, err = newHS("10♦", "2♦", "K♦", "5♦", "Q♥")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, HighCard, h.HandValue)
	Equal(t, "High card", h.EvalString())
}

func TestHandsCompare(t *testing.T) {

	// High card
	high, err := newHS("5♠", "3♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, HighCard, high.HandValue)

	// Pair
	hpair, err := newHS("A♠", "A♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Pair, hpair.HandValue)

	// Pair
	hpairlow, err := newHS("A♠", "2♥", "Q♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Pair, hpairlow.HandValue)

	// compare 2 pairs
	Equal(t, -1, hpair.Compare(hpairlow))

	// compare pair and high card
	Equal(t, -1, hpair.Compare(high))

	// Two pairs
	htwo, err := newHS("7♠", "7♥", "J♥", "A♥", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, TwoPairs, htwo.HandValue)

	// compare 2 pairs and high card
	Equal(t, -1, htwo.Compare(high))

	// compare 2 pairs and a pair
	Equal(t, -1, htwo.Compare(hpair))

	// Three
	hthree, err := newHS("A♠", "A♥", "J♥", "7♣", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, ThreeOfAKind, hthree.HandValue)

	// compare 3 and high card
	Equal(t, -1, hthree.Compare(high))

	// compare 3  and a pair
	Equal(t, -1, hthree.Compare(hpair))

	// compare 3  and 2 pairs
	Equal(t, -1, hthree.Compare(htwo))

}
