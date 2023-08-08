package singly

type (
	INode[T any] interface {
		Next() INode[T]
		Value() T
		SetNext(INode[T])
		SetValue(T)
	}

	Node[T any] struct {
		val  T
		next INode[T]
	}
)

func (n *Node[T]) SetNext(i INode[T]) {
	n.next = i
}

func (n *Node[T]) SetValue(t T) {
	n.val = t
}

func (n *Node[T]) Next() (next INode[T]) {
	if n == nil {
		return
	}
	return n.next
}

func (n *Node[T]) Value() (val T) {
	if n == nil {
		return
	}
	return n.val
}

var _ INode[int] = &Node[int]{}

func NewNode[T any](val T) INode[T] {
	return &Node[T]{val: val}
}
