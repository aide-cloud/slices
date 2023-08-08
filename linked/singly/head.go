package singly

import (
	"fmt"
	"strings"
)

type Head[T any] struct {
	// start 起始节点
	start INode[T]
	// length 链表长度
	length uint64
	// curr 当前节点
	curr INode[T]
	// end 结尾节点
	end INode[T]
}

type (
	IHead[T any] interface {
		Node() INode[T]
		Length() uint64
		IsEmpty() bool
		Clear()
		Curr() INode[T]
		Next() INode[T]
		// Append 尾插
		Append(val T) IHead[T]
		AppendNode(node INode[T]) IHead[T]
		// Prepend 头插
		Prepend(val T) IHead[T]
		PrependNode(node INode[T]) IHead[T]
		// Range 遍历
		Range(fn func(node INode[T]) bool)
		Slice() []T
		Show() string
		// Insert 插入
		Insert(index uint64, val T) IHead[T]
		InsertNode(index uint64, node INode[T]) IHead[T]
		// Remove 移除
		Remove(index uint64) IHead[T]
		RemoveValue(val T) IHead[T]
		RemoveHead() IHead[T]
		RemoveTail() IHead[T]
		// Reverse 反转
		Reverse() IHead[T]
		// Copy 复制
		Copy() IHead[T]
		// Find 查找
		Find(fn func(val T) bool) (INode[T], bool)
		FindIndex(fn func(val T) bool) (uint64, bool)
		FindAll(fn func(val T) bool) []INode[T]
	}
)

func New[T any](opts ...Option[T]) IHead[T] {
	head := &Head[T]{}
	for _, opt := range opts {
		opt(head)
	}
	return head
}

var _ IHead[int] = &Head[int]{} // 确保 Head[int] 实现了 IHead[int] 接口

func (h *Head[T]) FindAll(fn func(val T) bool) (nodes []INode[T]) {
	if h.IsEmpty() {
		return
	}
	nodes = make([]INode[T], 0, h.Length()/2)
	curr := h.start
	for curr != nil {
		if fn(curr.Value()) {
			nodes = append(nodes, curr)
		}
		curr = curr.Next()
	}
	return
}

func (h *Head[T]) FindIndex(fn func(val T) bool) (index uint64, ok bool) {
	if h.IsEmpty() {
		return
	}
	curr := h.start
	for curr != nil {
		if fn(curr.Value()) {
			return index, true
		}
		index++
		curr = curr.Next()
	}
	return
}

// Find 根据 val 查找节点
func (h *Head[T]) Find(fn func(val T) bool) (node INode[T], ok bool) {
	if h.IsEmpty() {
		return
	}
	curr := h.start
	for curr != nil {
		if fn(curr.Value()) {
			return curr, true
		}
		curr = curr.Next()
	}
	return
}

func (h *Head[T]) Copy() IHead[T] {
	head := New[T]()
	h.Range(func(node INode[T]) bool {
		head.Append(node.Value())
		return true
	})
	return head
}

func (h *Head[T]) Reverse() IHead[T] {
	if h.IsEmpty() {
		return h
	}
	var prev INode[T] = nil
	curr := h.start
	for curr != nil {
		next := curr.Next()
		curr.SetNext(prev)
		prev = curr
		curr = next
	}
	h.start = prev
	return h
}

func (h *Head[T]) Show() string {
	list := h.Slice()
	ss := make([]string, 0, len(list))
	for _, v := range list {
		ss = append(ss, fmt.Sprintf("(%+v)", v))
	}

	return fmt.Sprintf("%s", strings.Join(ss, "-->"))
}

func (h *Head[T]) Remove(index uint64) IHead[T] {
	if index >= h.length {
		return h
	}

	if index == h.length-1 {
		return h.RemoveTail()
	}

	if index == 0 {
		return h.RemoveHead()
	}

	prev := h.start
	for i := uint64(1); i < index; i++ {
		prev = prev.Next()
	}

	prev.SetNext(prev.Next().Next())
	h.length--
	return h
}

func (h *Head[T]) RemoveValue(val T) IHead[T] {
	panic("implement me")
}

func (h *Head[T]) RemoveTail() IHead[T] {
	if h.length == 0 {
		return h
	}

	if h.length == 1 {
		h.start = nil
		h.curr = nil
		h.end = nil
		h.length = 0
		return h
	}

	prev := h.start
	for i := uint64(1); i < h.length-1; i++ {
		prev = prev.Next()
	}

	prev.SetNext(nil)
	h.length--
	return h
}

func (h *Head[T]) RemoveHead() IHead[T] {
	if h.length == 0 {
		return h
	}

	h.start = h.start.Next()
	h.length--
	return h
}

func (h *Head[T]) Insert(index uint64, val T) IHead[T] {
	node := NewNode(val)
	return h.InsertNode(index, node)
}

func (h *Head[T]) InsertNode(index uint64, node INode[T]) IHead[T] {
	if index == 0 {
		return h.PrependNode(node)
	}

	if index >= h.length {
		return h.AppendNode(node)
	}

	prev := h.start
	for i := uint64(1); i < index; i++ {
		prev = prev.Next()
	}

	n := node
	h.length++
	for n.Next() != nil {
		h.length++
		n = n.Next()
	}
	n.SetNext(prev.Next())
	prev.SetNext(node)
	return h
}

func (h *Head[T]) Range(fn func(node INode[T]) bool) {
	for node := h.start; node != nil; node = node.Next() {
		if !fn(node) {
			break
		}
	}
}

func (h *Head[T]) Slice() []T {
	slice := make([]T, 0, h.length)
	for node := h.start; node != nil; node = node.Next() {
		slice = append(slice, node.Value())
	}
	return slice
}

func (h *Head[T]) AppendNode(node INode[T]) IHead[T] {
	h.length++
	if h.start == nil {
		h.setFirst(node)
		return h
	}

	h.end.SetNext(node)
	h.end = h.end.Next()
	return h
}

func (h *Head[T]) PrependNode(node INode[T]) IHead[T] {
	h.length++
	if h.start == nil {
		h.setFirst(node)
		return h
	}

	node.SetNext(h.start)
	h.start = node

	return h
}

func (h *Head[T]) setFirst(node INode[T]) {
	if h.start == nil {
		h.start = node
		h.end = h.start
	}
}

func (h *Head[T]) Append(val T) IHead[T] {
	h.length++
	if h.start == nil {
		h.setFirst(NewNode[T](val))
		return h
	}

	h.end.SetNext(NewNode[T](val))
	h.end = h.end.Next()
	return h
}

func (h *Head[T]) Prepend(val T) IHead[T] {
	h.length++
	if h.start == nil {
		h.setFirst(NewNode[T](val))
		return h
	}

	node := NewNode[T](val)
	node.SetNext(h.start)
	h.start = node

	return h
}

func (h *Head[T]) Clear() {
	h.start = nil
	h.length = 0
	h.curr = nil
}

func (h *Head[T]) Curr() INode[T] {
	if h.curr == nil {
		h.curr = h.start
	}
	return h.curr
}

func (h *Head[T]) Next() INode[T] {
	curr := h.Curr()
	if curr == nil {
		return nil
	}

	h.curr = curr.Next()
	return curr
}

func (h *Head[T]) Node() INode[T] {
	return h.start
}

func (h *Head[T]) Length() uint64 {
	return h.length
}

func (h *Head[T]) IsEmpty() bool {
	return h.length == 0
}