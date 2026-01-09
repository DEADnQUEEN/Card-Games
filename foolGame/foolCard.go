package foolGame

import (
	"cardGames/base"
	"cardGames/utils"
	"fmt"
	"github.com/fatih/color"
	"unicode/utf8"
)

const (
	Valet = 11
	Queen = 12
	King  = 13
	Ace   = 14

	Heart   = "\u2665"
	Diamond = "\u2666"
	Club    = "\u2663"
	Spades  = "\u2660"

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

type FoolCard struct {
	value int // Valet - 11, Queen - 12, King - 13, Ace - 14
	suit  int

	isMain bool
	isOpen bool
}

func CreateCard(value int, suit int, isMain bool) *FoolCard {
	return &FoolCard{
		value:  value,
		suit:   suit,
		isMain: isMain,
		isOpen: true,
	}
}

func (card *FoolCard) GetStringValue() string {
	var name, ok = valueNaming[card.value]
	if !ok {
		return fmt.Sprintf("%v", card.value)
	}
	if len(name) != 0 {
		return fmt.Sprintf("%c", name[0])
	}
	panic("name must have any letters")
}

func (card *FoolCard) GetStringSuit() string {
	if card.suit >= len(suits) {
		panic("out of range")
	}
	return suits[card.suit]
}

func (card *FoolCard) String() string {
	return card.GetStringValue() + card.GetStringSuit()
}

func (card *FoolCard) getStringsForClosedCard() []string {
	var strings = make([]string, 0, base.CardHeight)

	var line = ""
	for i := 2; i < base.CardWidth; i++ {
		line += base.HorizontalLine
	}
	strings = append(strings, base.TopLeftRound+line+base.TopRightRound)

	var patternLine = ""
	// 5 = border + border + padding inside 1 + padding inside 1 + 1 to add one more to end
	for j := 0; j < base.CardWidth-5; j++ {
		if j%2 != 0 {
			patternLine += BackPattern
		} else {
			patternLine += " "
		}
	}

	for i := 2; i < base.CardHeight; i++ {
		var heightLine = base.VerticalLine + " "

		if i%2 == 0 {
			heightLine += color.RedString(BackPattern + patternLine)
		} else {
			heightLine += color.RedString(patternLine + " ")
		}

		strings = append(strings, heightLine+" "+base.VerticalLine)
	}

	return append(strings, base.BottomLeftRound+line+base.BottomRightRound)
}

func (card *FoolCard) getStringsForOpenCard() []string {
	var strings = make([]string, 0, base.CardHeight)

	var value = card.GetStringValue()
	var suit = card.GetStringSuit()
	var text = " " + value + " " + suit + " "
	var coloredText = colorTextForSuit(text, suit)
	if card.value != 10 {
		text = base.HorizontalLine + text
		coloredText = base.HorizontalLine + coloredText
	}

	var line = ""
	for i := utf8.RuneCountInString(text) + 2; i < base.CardWidth; i++ {
		line += base.HorizontalLine
	}
	strings = append(strings, base.TopLeftRound+line+coloredText+base.TopRightRound)

	var heightLine = base.VerticalLine
	for i := 2; i < base.CardWidth; i++ {
		heightLine += " "
	}

	for i := 1; i < base.CardHeight-1; i++ {
		strings = append(strings, heightLine+base.VerticalLine)
	}

	return append(strings, base.BottomLeftRound+coloredText+line+base.BottomRightRound)

}

func (card *FoolCard) GetStrings() []string {
	if card.isOpen {
		return card.getStringsForOpenCard()
	}
	return card.getStringsForClosedCard()
}

func (card *FoolCard) ShowCard() string {
	var cardStrings = card.GetStrings()
	var show = ""
	for _, cardString := range cardStrings {
		show += cardString + "\n"
	}

	return show
}

func (card *FoolCard) CompareTo(to utils.Comparable) int {
	switch to.(type) {
	case *FoolCard:
		var compareCard = to.(*FoolCard)
		if compareCard.isMain && !card.isMain {
			return utils.Bigger
		}
		if !compareCard.isMain && card.isMain {
			return utils.Less
		}
		if compareCard.suit != card.suit {
			return utils.Error
		}
		if compareCard.value > card.value {
			return utils.Bigger
		}
		if compareCard.value < card.value {
			return utils.Less
		}

		return utils.Equal
	}

	panic("couldn't compare FoolCard")
}
