package base

type Player interface {
	Action() error
	GetInfo() string
	IsPlaying() bool
}
