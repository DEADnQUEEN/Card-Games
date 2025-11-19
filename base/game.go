package base

import (
	"cardGames/utils"
	"errors"
	"fmt"
)

var cardsCount = map[int]int{
	36: 6,
	52: 2,
}

type FoolGame struct {
	playableCards []*DefaultCard

	deck    utils.Queue[*DefaultCard]
	players []Player
}

func (game *FoolGame) ShowCards() {
	var width = len(fmt.Sprintf("%d", len(game.playableCards)))
	for index, card := range game.playableCards {
		var skip = ""

		for i := len(fmt.Sprintf("%d", index)); i < width; i++ {
			skip += " "
		}

		fmt.Printf("%s%d | %s\n", skip, index, card)
	}
}

func (game *FoolGame) ShuffleDeck() {

}

// NewFoolGame Cards Amount is 36, 52
func NewFoolGame(cardsAmount int, playerCount int) (*FoolGame, error) {
	var start, ok = cardsCount[cardsAmount]
	if !ok {
		return nil, errors.New("cardsAmount must be equal to 36 or 52")
	}

	var cards = make([]*DefaultCard, cardsAmount)
	var deck = utils.Queue[*DefaultCard]{}

	cardsAmount /= len(suits)

	var mainSuit = 1 // random

	for cardValue := 0; cardValue < cardsAmount; cardValue++ {
		for suit := 0; suit < len(suits); suit++ {
			var card = &DefaultCard{
				value:  cardValue + start,
				suit:   suit,
				isMain: suit == mainSuit,
			}
			cards[cardValue+(cardsAmount*suit)] = card
			deck.Enqueue(card)
		}
	}

	return &FoolGame{
		playableCards: cards,
		deck:          deck,
		players:       make([]Player, playerCount),
	}, nil
}
