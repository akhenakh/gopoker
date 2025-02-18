package poker

import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

var suits [4]string = [4]string{"♦", "♠", "♥", "♣"}

// Card a card representation
type Card struct {
	// value as 1 = A♦ , 2 = 2♦, 11 = J♦, 12 = Q♦, 13 = K♦, 14 = A♠
	StoreValue int
}

// New create a new Card checking passed values
func NewCard(value int) (*Card, error) {
	if value < 1 || value > 52 {
		return nil, errors.New("card value out of limits")
	}
	return &Card{value}, nil
}

// newCS create a new card bypassing err and with a string as parameters
// main purpose is to simplifiy testing
// v is a string as follow "A♥"
func newCS(v string) *Card {
	if len(v) < 2 {
		return nil
	}
	var value int
	var suit int

	switch v[0] {
	case 'J':
		value = 11
	case 'Q':
		value = 12
	case 'K':
		value = 13
	case 'A':
		value = 1
	case '2', '3', '4', '5', '6', '7', '8', '9':
		vint, _ := strconv.Atoi(string(v[0]))
		value = int(vint)
	}
	if v[0:2] == "10" {
		value = 10
	}

	r, _ := utf8.DecodeLastRuneInString(v)
	switch r {
	case '♦':
		suit = 0
	case '♠':
		suit = 1
	case '♥':
		suit = 2
	case '♣':
		suit = 3
	}
	c, _ := NewCard(value + (suit * 13))
	return c
}

// Value return the value of the card A = 14, 10 = 10, Q = 11
func (c *Card) Value() int {
	// Ace case
	if c.StoreValue%13 == 0 {
		return 13
	}
	if c.StoreValue%13 == 1 {
		return 14
	}
	return c.StoreValue % 13
}

// Suit return the suit as follow 0 = ♦, 1 = ♠, 2 = ♥, 3 = ♣
func (c *Card) Suit() int {
	s := c.StoreValue / 13
	n := c.StoreValue % 13
	if n == 0 {
		s--
	}
	return s
}

// valid return true if a card is a valid one
func (c *Card) valid() bool {
	if c.StoreValue < 1 || c.StoreValue > 52 {
		return false
	}
	return true
}

// String display a Card as string
func (c *Card) String() string {
	if !c.valid() {
		return "Undefined"
	}
	return fmt.Sprintf("%s%s", c.ValueAsString(), c.SuitAsString())
}

// valueAsString returns the poker value of the Card
func (c *Card) ValueAsString() string {
	return cardValueAsString(c.Value())
}

func cardValueAsString(v int) string {
	switch v {
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	case 2, 3, 4, 5, 6, 7, 8, 9, 10:
		return strconv.Itoa(int(v))
	}
	return "Undefined"
}

// suiteAsString returns the poker suit of the Card
func (c *Card) SuitAsString() string {
	if c.Suit() >= 0 && c.Suit() <= 3 {
		return suits[c.Suit()]
	}
	return "Undefined"
}

type CardSlice []Card

func (cs CardSlice) Len() int           { return len(cs) }
func (cs CardSlice) Less(i, j int) bool { return cs[i].Value() < cs[j].Value() }
func (cs CardSlice) Swap(i, j int)      { cs[i], cs[j] = cs[j], cs[i] }
