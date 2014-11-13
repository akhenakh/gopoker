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
	// value as 2 = 2, 11 = J, 12 = Q, 13 = K, 14 = A
	value int
	// suit as 1 = ♦, 2 = ♠, 3 = ♥, 4 = ♣
	suit int
}

// New create a new Card checking passed values
func NewCard(value, suit int) (*Card, error) {
	if value < 2 || value > 14 {
		return nil, errors.New("card value out of limits")
	}
	if suit < 1 || suit > 4 {
		return nil, errors.New("card suit out of limits")
	}
	return &Card{value, suit}, nil
}

// newCardSuit create a new card bypassing err and with strings as parameters
// main purpose is to simplifiy testing
// v is a string as follow A♥
func newCardString(v string) *Card {
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
		value = 14
	case '2', '3', '4', '5', '6', '7', '8', '9':
		value, _ = strconv.Atoi(string(v[0]))
	}
	if v[0:2] == "10" {
		value = 10
	}

	r, _ := utf8.DecodeLastRuneInString(v)
	switch r {
	case '♦':
		suit = 1
	case '♠':
		suit = 2
	case '♥':
		suit = 3
	case '♣':
		suit = 4
	}

	c, _ := NewCard(value, suit)
	return c
}

// String display a Card as string
func (c *Card) String() string {
	return fmt.Sprintf("%s%s", c.ValueAsString(), c.SuitAsString())
}

// valueAsString returns the poker value of the Card
func (c *Card) ValueAsString() string {
	switch c.value {
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	case 14:
		return "A"
	case 2, 3, 4, 5, 6, 7, 8, 9, 10:
		return strconv.Itoa(c.value)
	}
	return "Undefined"
}

// suiteAsString returns the poker suit of the Card
func (c *Card) SuitAsString() string {
	if c.suit > 0 && c.suit < 5 {
		return suits[c.suit-1]
	}
	return "Undefined"
}
