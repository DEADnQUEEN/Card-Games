package utils

const (
	Less   = 0
	Equal  = 1
	Bigger = 2
	Error  = 3
)

type Comparable interface {
	CompareTo(item Comparable) int
}
