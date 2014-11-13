package poker

import (
	"errors"
	"fmt"
	"strconv"
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
func NewCard(value, suit int) (error, *Card) {
	if value < 2 || value > 14 {
		return errors.New("card value out of limits"), nil
	}
	if suit < 1 || suit > 4 {
		return errors.New("card suit out of limits"), nil
	}
	return nil, &Card{value, suit}
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
