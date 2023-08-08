package singly

type Option[T any] func(*Head[T])

func WithNodes[T any](nodes ...INode[T]) Option[T] {
	return func(h *Head[T]) {
		for _, node := range nodes {
			h.AppendNode(node)
		}
	}
}

func WithValues[T any](values ...T) Option[T] {
	return func(h *Head[T]) {
		for _, val := range values {
			h.Append(val)
		}
	}
}
