package foolGame

import (
	"cardGames/base"
	"cardGames/utils"
	"errors"
	"fmt"
)

var cardsCount = map[int]int{
	36: 6,
	52: 2,
}

type FoolGame struct {
	playableCards []*FoolCard

	deck    utils.Queue[*FoolCard]
	players []FoolPlayer
}

func (game *FoolGame) ShowAllPlayableCards() {
	fmt.Println("Playable cards: ")

	var group = len(game.playableCards) / len(suits)

	for i := 0; i < len(game.playableCards); i += group {
		var cards = make([]base.Card, 0, group)
		for j := 0; j < group && j+i < len(game.playableCards); j++ {
			cards = append(cards, game.playableCards[i+j])
		}
		fmt.Println(base.StackCardsInRow(cards))
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

	var cards = make([]*FoolCard, cardsAmount)
	var deck = utils.Queue[*FoolCard]{}

	cardsAmount /= len(suits)

	var mainSuit = 1 // random

	for cardValue := 0; cardValue < cardsAmount; cardValue++ {
		for suit := 0; suit < len(suits); suit++ {
			var card = CreateCard(cardValue+start, suit, suit == mainSuit)
			cards[cardValue+(cardsAmount*suit)] = card
			deck.Enqueue(card)
		}
	}

	return &FoolGame{
		playableCards: cards,
		deck:          deck,
		players:       make([]FoolPlayer, playerCount),
	}, nil
}
