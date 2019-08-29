package bptree

const (
	MaxKV = 255
	maxKC = 511
)

type node interface {
	find(key int) (int, bool)
	parent() *interiorNode
	setParent(*interiorNode)
	full() bool
}