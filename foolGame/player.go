package foolGame

import (
	"cardGames/base"
	"fmt"
)

type FoolPlayer struct {
	playerCards []base.Card
	name        string
}

func (f *FoolPlayer) ShowCards() {
	fmt.Println(base.StackCardsInRow(f.playerCards))
}

func (f *FoolPlayer) Action() error {
	//TODO implement me
	panic("implement me")
}

func (f *FoolPlayer) GetInfo() string {
	return fmt.Sprintf("PLayer \"%s\", cards amount: %d", f.name, len(f.playerCards))
}

func (f *FoolPlayer) IsPlaying() bool {
	//TODO implement me
	panic("implement me")
}
