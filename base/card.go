package base

import (
	"cardGames/utils"
	"fmt"
	"github.com/fatih/color"
	"unicode/utf8"
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

	TopLeftRound     = "\u256D"
	TopRightRound    = "\u256E"
	BottomLeftRound  = "\u2570"
	BottomRightRound = "\u256F"

	HorizontalLine = "\u2500"
	VerticalLine   = "\u2502"

	CardWidth   = 11
	CardHeight  = 7
	BackPattern = "+"
)

var suits = []string{Heart, Diamond, Club, Spades}
var valueNaming = map[int]string{
	Valet: "Valet",
	Queen: "Queen",
	King:  "King",
	Ace:   "Ace",
}

func colorTextForSuit(text string, suit string) string {
	switch suit {
	case Heart:
		return color.RedString(text)
	case Diamond:
		return color.RedString(text)
	case Club:
		return color.BlueString(text)
	case Spades:
		return color.BlueString(text)
	}
	return text
}

type DefaultCard struct {
	value int // Valet - 11, Queen - 12, King - 13, Ace - 14
	suit  int

	isMain bool
	isOpen bool
}

func CreateCard(value int, suit int, isMain bool) *DefaultCard {
	return &DefaultCard{
		value:  value,
		suit:   suit,
		isMain: isMain,
		isOpen: true,
	}
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
	return card.GetStringValue() + card.GetStringSuit()
}

func (card *DefaultCard) getStringsForClosedCard() []string {
	var strings = make([]string, 0, CardHeight)

	var line = ""
	for i := 2; i < CardWidth; i++ {
		line += HorizontalLine
	}
	strings = append(strings, TopLeftRound+line+TopRightRound)

	var patternLine = ""
	// 5 = border + border + padding inside 1 + padding inside 1 + 1 to add one more to end
	for j := 0; j < CardWidth-5; j++ {
		if j%2 != 0 {
			patternLine += BackPattern
		} else {
			patternLine += " "
		}
	}

	for i := 2; i < CardHeight; i++ {
		var heightLine = VerticalLine + " "

		if i%2 == 0 {
			heightLine += color.RedString(BackPattern + patternLine)
		} else {
			heightLine += color.RedString(patternLine + " ")
		}

		strings = append(strings, heightLine+" "+VerticalLine)
	}

	return append(strings, BottomLeftRound+line+BottomRightRound)
}

func (card *DefaultCard) getStringsForOpenCard() []string {
	var strings = make([]string, 0, CardHeight)

	var value = card.GetStringValue()
	var suit = card.GetStringSuit()
	var text = " " + value + " " + suit + " "
	var coloredText = colorTextForSuit(text, suit)
	if card.value != 10 {
		text = HorizontalLine + text
		coloredText = HorizontalLine + coloredText
	}

	var line = ""
	for i := utf8.RuneCountInString(text) + 2; i < CardWidth; i++ {
		line += HorizontalLine
	}
	strings = append(strings, TopLeftRound+line+coloredText+TopRightRound)

	var heightLine = VerticalLine
	for i := 2; i < CardWidth; i++ {
		heightLine += " "
	}

	for i := 1; i < CardHeight-1; i++ {
		strings = append(strings, heightLine+VerticalLine)
	}

	return append(strings, BottomLeftRound+coloredText+line+BottomRightRound)

}

func (card *DefaultCard) getStrings() []string {
	if card.isOpen {
		return card.getStringsForOpenCard()
	}
	return card.getStringsForClosedCard()
}

func (card *DefaultCard) ShowCard() string {
	var cardStrings = card.getStrings()
	var show = ""
	for _, cardString := range cardStrings {
		show += cardString + "\n"
	}

	return show

}

func StackCardsInRow(cards []*DefaultCard) string {
	var cardsView = make([][]string, len(cards))
	for i, card := range cards {
		cardsView[i] = card.getStrings()
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
