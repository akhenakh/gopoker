package poker

import (
	"fmt"
	"sort"
)

type HandValue int

const (
	_ HandValue = iota
	HighCard
	Pair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

// Hand is a 5 Cards poker hand
type Hand struct {
	Cards     [5]Card
	HandValue HandValue
	value     [5]int
}

// NewHand is returning a *Hand and evaluate it
func NewHand(cards [5]Card) (*Hand, error) {
	for i, c := range cards {
		if !c.valid() {
			return nil, fmt.Errorf("Invalid card at pos %d", i)
		}
	}

	// checking for duplicates
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				continue
			}
			if cards[i].value == cards[j].value {
				return nil, fmt.Errorf("Invalid hand duplicates on %d %d", i, j)
			}
		}

	}
	h := &Hand{Cards: cards}
	h.evaluate()
	return h, nil
}

// newHS new hand as Strings, mainly to simplify testing
// returning a *Hand and evaluate it
func newHS(sc ...string) (*Hand, error) {

	if len(sc) != 5 {
		return nil, fmt.Errorf("Invalid number of cards")
	}
	var cards [5]Card
	for i, s := range sc {
		c := newCS(s)
		if c == nil {
			return nil, fmt.Errorf("Invalid card at pos %d", i)
		}
		cards[i] = *c
	}

	return NewHand(cards)
}

// evaluate is evaluating the poker hand and store results in value
func (h *Hand) evaluate() {
	var flush bool
	var straightLow int

	// we are looking for pairs, double pairs, threes, fours
	// map groupped hands
	mgh := make(map[int]int, 0)

	// grouping by []int of the same card value
	for _, c := range h.Cards {
		v := c.Value()
		mgh[v]++
	}
	sc := make([]int, 0)
	gh := make([][]int, 0)

	for k, v := range mgh {
		if v == 1 {
			// append to solocard
			sc = append(sc, k)
		} else {
			// grouping by same value
			l := make([]int, 0)
			for i := int(1); i <= v; i++ {
				l = append(l, k)
			}
			gh = append(gh, l)
		}
	}

	if len(gh) > 0 {
		// sort the remaining solo card
		sort.Sort(sort.IntSlice(sc))
		// inverse map remaining solo card to value
		for x := 0; x < len(sc); x++ {
			h.value[4-x] = sc[x]
		}

		if len(gh) == 1 {
			if len(gh[0]) == 2 {
				// we got a pair
				h.HandValue = Pair
				h.value[0], h.value[1] = gh[0][0], gh[0][0]
			} else if len(gh[0]) == 3 {
				// we got three
				h.HandValue = ThreeOfAKind
				h.value[0], h.value[1], h.value[2] = gh[0][0], gh[0][0], gh[0][0]
			} else if len(gh[0]) == 4 {
				// we got four
				h.HandValue = FourOfAKind
				h.value[0], h.value[1], h.value[2], h.value[3] = gh[0][0], gh[0][0], gh[0][0], gh[0][0]
			}
		} else if len(gh) == 2 {
			// 2 paires
			if len(gh[0]) == 2 && len(gh[1]) == 2 {
				h.HandValue = TwoPairs
				// order the 2 pairs higher first
				if gh[0][0] < gh[1][0] {
					gh[0], gh[1] = gh[1], gh[0]
				}
				h.value[0], h.value[1] = gh[0][0], gh[0][0]
				h.value[2], h.value[3] = gh[1][0], gh[1][0]
			} else {
				// full house
				h.HandValue = FullHouse
				// order the full house the triple first
				if len(gh[0]) < len(gh[1]) {
					gh[0], gh[1] = gh[1], gh[0]
				}
				h.value[0], h.value[1], h.value[2] = gh[0][0], gh[0][0], gh[0][0]
				h.value[3], h.value[4] = gh[1][0], gh[1][0]
			}
		}
		return
	}

	// copy the cards and sort by value
	cards := make(CardSlice, 5)
	for i, c := range h.Cards {
		cards[i] = c
	}
	sort.Sort(cards)

	for i := 0; i < 5; i++ {
		h.value[4-i] = cards[i].Value()
	}

	// test for flush or straight
	if cards[0].Suit() == cards[1].Suit() && cards[2].Suit() == cards[3].Suit() &&
		cards[4].Suit() == cards[0].Suit() && cards[0].Suit() == cards[2].Suit() {
		flush = true
	}

	// test low Straight for Ace
	if cards[0].Value() == 2 && cards[1].Value() == 3 && cards[2].Value() == 4 &&
		cards[3].Value() == 5 && cards[4].Value() == 14 {
		straightLow = 1
	}

	// test other straight
	for x := 2; x < 11; x++ {
		if cards[0].Value() == x && cards[1].Value() == x+1 && cards[2].Value() == x+2 &&
			cards[3].Value() == x+3 && cards[4].Value() == x+4 {
			straightLow = x
			break
		}
	}

	if straightLow > 0 {
		for i := 0; i < 5; i++ {
			h.value[4-i] = straightLow + i
		}
	}

	if straightLow > 0 && !flush {
		h.HandValue = Straight
	} else if straightLow > 0 && flush {
		h.HandValue = StraightFlush
	} else if flush {
		h.HandValue = Flush
	}

	if h.HandValue == 0 {
		h.HandValue = HighCard
	}
}

// Compare will compare two hands returning -1 when oh is lower, 1 when oh is higher, 0 if deuce
func (h *Hand) Compare(oh *Hand) int {
	if oh.HandValue > h.HandValue {
		return 1
	} else if oh.HandValue < h.HandValue {
		return -1
	}

	for i := 0; i < 5; i++ {
		if oh.value[i] > h.value[i] {
			return 1
		} else if oh.value[i] < h.value[i] {
			return -1
		}
	}
	return 0
}

// String display a Card as string
func (h *Hand) String() string {
	return fmt.Sprintf("%s %s %s %s %s",
		h.Cards[0].String(), h.Cards[1].String(), h.Cards[2].String(), h.Cards[3].String(), h.Cards[4].String())
}

// EvalString display the poker evaluated hand as string
func (h *Hand) EvalString() string {
	switch HandValue(h.HandValue) {
	case HighCard:
		return "High card"
	case Pair:
		return fmt.Sprintf("Pair of %s's", cardValueAsString(h.value[0]))
	case TwoPairs:
		return fmt.Sprintf("Two pair %s %s", cardValueAsString(h.value[0]), cardValueAsString(h.value[2]))
	case ThreeOfAKind:
		return fmt.Sprintf("Three of a kind %s's", cardValueAsString(h.value[0]))
	case Flush:
		return fmt.Sprintf("%s high flush", cardValueAsString(h.value[0]))
	case Straight:
		return fmt.Sprintf("Straight %s high", cardValueAsString(h.value[0]))
	case FullHouse:
		return fmt.Sprintf("Full house %s over %s", cardValueAsString(h.value[0]), cardValueAsString(h.value[3]))
	case FourOfAKind:
		return fmt.Sprintf("Four of a kind %s's", cardValueAsString(h.value[0]))
	case StraightFlush:
		return fmt.Sprintf("Straight flush %s high", cardValueAsString(h.value[0]))
	}
	return "Invalid hand"
}
