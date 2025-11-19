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
)

var suits = []string{"Heart", "Diamond", "Club", "Spades"}

var naming = map[int]string{
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
	var name, ok = naming[card.value]
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
	var suit = suits[card.suit]
	if len(suit) == 0 {
		panic("suit must have any letters")
	}

	return fmt.Sprintf("%c", suit[0])
}

func (card *DefaultCard) String() string {
	var value = card.GetStringValue()
	if len(value) == 1 {
		value = " " + value
	}
	return value + card.GetStringSuit()
}
