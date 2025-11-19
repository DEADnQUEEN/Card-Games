package base

import (
	"cardGames/utils"
	"fmt"
)

type Card interface {
	utils.Comparable
}

const (
	Valet = 11
	Queen = 12
	King  = 13
	Ace   = 14

	Heart   = "\u2665"
	Diamond = "\u2666"
	Club    = "\u2663"
	Spades  = "\u2660"
)

var suits = []string{Heart, Diamond, Club, Spades}

var valueNaming = map[int]string{
	Valet: "Valet",
	Queen: "Queen",
	King:  "King",
	Ace:   "Ace",
}

type DefaultCard struct {
	value int // Valet - 11, Queen - 12, King - 13, Ace - 14
	suit  int

	isMain bool
}

func (card *DefaultCard) GetStringValue() string {
	var name, ok = valueNaming[card.value]
	if !ok {
		return fmt.Sprintf("%v", card.value)
	}
	if len(name) != 0 {
		return fmt.Sprintf("%c", name[0])
	}
	panic("name must have any letters")
}

func (card *DefaultCard) GetStringSuit() string {
	if card.suit >= len(suits) {
		panic("out of range")
	}
	return suits[card.suit]
}

func (card *DefaultCard) String() string {
	var value = card.GetStringValue()
	if len(value) == 1 {
		value = " " + value
	}
	return value + card.GetStringSuit()
}
