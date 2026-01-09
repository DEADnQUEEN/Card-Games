package base

import (
	"cardGames/utils"
)

const (
	TopLeftRound     = "\u256D"
	TopRightRound    = "\u256E"
	BottomLeftRound  = "\u2570"
	BottomRightRound = "\u256F"

	HorizontalLine = "\u2500"
	VerticalLine   = "\u2502"

	CardWidth  = 14
	CardHeight = 12
)

type Card interface {
	utils.Comparable

	GetStrings() []string
	ShowCard() string
}

func StackCardsInRow(cards []Card) string {
	var cardsView = make([][]string, len(cards))
	for i, card := range cards {
		cardsView[i] = card.GetStrings()
	}

	var connected = ""
	for i := 0; i < CardHeight; i++ {
		for j, card := range cardsView {
			if j != 0 {
				connected += "   "
			}
			connected += card[i]
		}
		connected += "\n"
	}

	return connected
}
