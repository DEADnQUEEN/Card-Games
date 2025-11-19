package base

type Player interface {
	Action() error
	GetInfo() error
	isPlaying() bool
}
