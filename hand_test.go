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

	// Straight
	hstraight, err := newHS("A♣", "2♣", "3♦", "4♦", "5♥")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Straight, hstraight.HandValue)
	// compare straight and high card
	Equal(t, -1, hstraight.Compare(high))
	// compare straight  and a pair
	Equal(t, -1, hstraight.Compare(hpair))
	// compare straight  and 2 pairs
	Equal(t, -1, hstraight.Compare(htwo))
	// compare straight  and threes
	Equal(t, -1, hstraight.Compare(hthree))

	// Flush
	hflush, err := newHS("7♥", "9♥", "3♥", "4♥", "2♥")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Flush, hflush.HandValue)
	// compare flush and high card
	Equal(t, -1, hflush.Compare(high))
	// compare flush  and a pair
	Equal(t, -1, hflush.Compare(hpair))
	// compare flush  and 2 pairs
	Equal(t, -1, hflush.Compare(htwo))
	// compare flush  and threes
	Equal(t, -1, hflush.Compare(hthree))
	// compare flush  and straight
	Equal(t, -1, hflush.Compare(hstraight))

	// Full House
	hfull, err := newHS("A♦", "A♥", "J♥", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, FullHouse, hfull.HandValue)
	// compare full and high card
	Equal(t, -1, hfull.Compare(high))
	// compare full  and a pair
	Equal(t, -1, hfull.Compare(hpair))
	// compare full  and 2 pairs
	Equal(t, -1, hfull.Compare(htwo))
	// compare full  and threes
	Equal(t, -1, hfull.Compare(hthree))
	// compare full  and straight
	Equal(t, -1, hfull.Compare(hstraight))
	// compare full  and flush
	Equal(t, -1, hfull.Compare(hflush))

	// Four
	hfour, err := newHS("A♦", "A♥", "J♥", "A♠", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, FourOfAKind, hfour.HandValue)
	// compare four and high card
	Equal(t, -1, hfour.Compare(high))
	// compare four  and a pair
	Equal(t, -1, hfour.Compare(hpair))
	// compare four  and 2 pairs
	Equal(t, -1, hfour.Compare(htwo))
	// compare four  and threes
	Equal(t, -1, hfour.Compare(hthree))
	// compare four  and straight
	Equal(t, -1, hfour.Compare(hstraight))
	// compare four  and flush
	Equal(t, -1, hfour.Compare(hflush))
	// compare four  and full
	Equal(t, -1, hfour.Compare(hfull))

	// Straight flush
	hsf, err := newHS("A♦", "2♦", "3♦", "4♦", "5♦")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, StraightFlush, hsf.HandValue)
	// compare straight flush and high card
	Equal(t, -1, hsf.Compare(high))
	Equal(t, 1, high.Compare(hsf))
	// compare straight flush  and a pair
	Equal(t, -1, hsf.Compare(hpair))
	Equal(t, 1, hpair.Compare(hsf))
	// compare straight flush  and 2 pairs
	Equal(t, -1, hsf.Compare(htwo))
	Equal(t, 1, htwo.Compare(hsf))
	// compare straight flush  and threes
	Equal(t, -1, hsf.Compare(hthree))
	Equal(t, 1, hthree.Compare(hsf))
	// compare straight flush  and straight
	Equal(t, -1, hsf.Compare(hstraight))
	Equal(t, 1, hstraight.Compare(hsf))
	// compare straight flush  and flush
	Equal(t, -1, hsf.Compare(hflush))
	Equal(t, 1, hflush.Compare(hsf))
	// compare straight flush  and full
	Equal(t, -1, hsf.Compare(hfull))
	Equal(t, 1, hfull.Compare(hsf))
	// compare straight flush  and four
	Equal(t, -1, hsf.Compare(hfour))
	Equal(t, 1, hfour.Compare(hsf))
}

func TestBestHands(t *testing.T) {
	// High card
	high, err := newBestHS("5♠", "3♥", "J♥", "7♣", "Q♠", "2♠", "8♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, HighCard, high.HandValue)

	// Pair
	h, err := newBestHS("3♦", "5♦", "A♠", "A♥", "J♥", "7♣", "Q♠")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, Pair, h.HandValue)
	Equal(t, "Pair of A's", h.EvalString())

	// Two pairs
	h, err = newBestHS("7♠", "3♦", "5♦", "7♥", "J♥", "A♥", "A♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, TwoPairs, h.HandValue)
	Equal(t, "Two pair A 7", h.EvalString())

	// Straight flush
	hsf, err := newBestHS("A♦", "2♦", "3♦", "4♦", "5♦", "J♠", "J♣")
	if err != nil {
		t.Fatal(err)
	}
	Equal(t, StraightFlush, hsf.HandValue)

}
