package singly

import (
	"reflect"
	"testing"
)

func TestHead_AppendNode(t *testing.T) {
	type args[T any] struct {
		node INode[T]
	}
	type testCase[T any] struct {
		name string
		h    Head[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "empty",
			h: Head[int]{
				start:  nil,
				length: 0,
				curr:   nil,
			},
			args: args[int]{
				node: NewNode(1),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.AppendNode(tt.args.node); !reflect.DeepEqual(got.Curr().Value(), tt.want) {
				t.Errorf("AppendNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHead_Range(t *testing.T) {
	type args[T any] struct {
		fn func(node INode[T]) bool
	}
	type testCase[T any] struct {
		name string
		h    Head[T]
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "empty",
			h: Head[int]{
				start:  nil,
				length: 0,
				curr:   nil,
			},
			args: args[int]{
				fn: func(node INode[int]) bool {
					return false
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Range(tt.args.fn)
		})
	}
}

func TestHead_Slice(t *testing.T) {
	type testCase[T any] struct {
		name string
		h    IHead[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "empty",
			h:    New[int](),
			want: []int{},
		},
		{
			name: "one",
			h:    New[int](WithNodes[int](NewNode(1))),
			want: []int{1},
		},
		{
			name: "two",
			h:    New[int](WithNodes[int](NewNode(1), NewNode(2))),
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHead_AppendNode1(t *testing.T) {
	h := New[int]()
	for i := 0; i < 10; i++ {
		h.AppendNode(NewNode(i))
	}

	if reflect.DeepEqual(h.Slice(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_AppendNode2(t *testing.T) {
	h := New[int](WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}
	h.Append(11).AppendNode(NewNode(12))
	if h.Length() == 12 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_InsertNode(t *testing.T) {
	h := New[int]()

	h.InsertNode(0, NewNode(1)).Insert(0, 2).Insert(0, 3)
	if h.Length() == 3 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{3, 2, 1}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	node := NewNode(4)
	node.SetNext(NewNode(5))
	h.InsertNode(1, node)
	if h.Length() == 5 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{3, 4, 5, 2, 1}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_RemoveHead(t *testing.T) {
	h := New[int](WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	h.RemoveHead()
	if h.Length() == 9 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_Remove(t *testing.T) {
	h := New[int](WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	h.Remove(0).Remove(0).Remove(0)
	if h.Length() == 7 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{4, 5, 6, 7, 8, 9, 10}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_RemoveTail(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	h.RemoveTail().RemoveTail().RemoveTail()
	if h.Length() == 7 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h.Slice(), []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_Show(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	t.Log(h.Show())

	type User struct {
		Name string
		Age  int
	}

	h2 := New(WithValues(User{"张三", 18}, User{"李四", 19}, User{"王五", 20}))
	t.Log(h2.Show())
	t.Log(h2.Slice())
}

func TestHead_Reverse(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	t.Log(h.Show())

	h.Reverse()
	t.Log(h.Show())
	h.Reverse()
	t.Log(h.Show())
}

func TestHead_Copy(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	t.Log(h.Show())

	h2 := h.Copy()
	t.Log(h2.Show())
	h2.Append(11)
	t.Log(h2.Show())
	t.Log(h.Show())
}

func TestHead_ReverseCopy(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	h2 := h.Reverse().Copy()
	if h2.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	if reflect.DeepEqual(h2.Slice(), []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}) {
		t.Log("ok")
	} else {
		t.Error("err")
	}
	t.Log(h2.Show())
}

func TestHead_Find(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	node, ok := h.Find(func(val int) bool {
		return val == 5
	})
	if ok {
		t.Log(node.Value())
		if node.Value() == 5 {
			t.Log("ok")
		} else {
			t.Error("err")
		}
	} else {
		t.Error("err")
	}

	node, ok = h.Find(func(val int) bool {
		return val == 11
	})
	if !ok {
		t.Log("ok")
	} else {
		t.Error("err")
	}
}

func TestHead_FindIndex(t *testing.T) {
	type args[T any] struct {
		fn func(val T) bool
	}
	type testCase[T any] struct {
		name      string
		h         IHead[T]
		args      args[T]
		wantIndex uint64
		wantOk    bool
	}
	tests := []testCase[int]{
		{
			name:      "1",
			h:         New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)),
			args:      args[int]{fn: func(val int) bool { return val == 5 }},
			wantIndex: 4,
			wantOk:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOk := tt.h.FindIndex(tt.args.fn)
			if gotIndex != tt.wantIndex {
				t.Errorf("FindIndex() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FindIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestHead_Find1(t *testing.T) {
	type args[T any] struct {
		fn func(val T) bool
	}
	type testCase[T any] struct {
		name        string
		h           IHead[T]
		args        args[T]
		wantNodeVal T
		wantOk      bool
	}
	tests := []testCase[int]{
		{
			name:        "1",
			h:           New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)),
			args:        args[int]{fn: func(val int) bool { return val == 5 }},
			wantNodeVal: 5,
			wantOk:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNode, gotOk := tt.h.Find(tt.args.fn)
			if !reflect.DeepEqual(gotNode.Value(), tt.wantNodeVal) {
				t.Errorf("Find() gotNode = %v, want %v", gotNode.Value(), tt.wantNodeVal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Find() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestHead_FindAll(t *testing.T) {
	type args[T any] struct {
		fn func(val T) bool
	}
	type testCase[T any] struct {
		name           string
		h              IHead[T]
		args           args[T]
		wantNodeValues []T
	}
	tests := []testCase[int]{
		{
			name:           "1",
			h:              New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)),
			args:           args[int]{fn: func(val int) bool { return val > 5 }},
			wantNodeValues: []int{6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNodes := tt.h.FindAll(tt.args.fn)
			for i, gotNode := range gotNodes {
				if !reflect.DeepEqual(gotNode.Value(), tt.wantNodeValues[i]) {
					t.Errorf("FindAll() gotNode = %v, want %v", gotNode.Value(), tt.wantNodeValues[i])
				}
			}
		})
	}
}

func TestHead_Clear(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	t.Log(h.Show())

	h.Clear()

	if h.Length() == 0 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	t.Log(h.Show())
}

func TestHead_Curr(t *testing.T) {
	h := New(WithValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	for i := 0; i < 10; i++ {
		if h.Curr().Value() == i+1 {
			t.Log("ok")
		} else {
			t.Error("err")
		}
		h.Next()
	}
}

func TestHead_Sort(t *testing.T) {
	h := New(WithValues(1, 2, 4, 3, 5, 6, 7, 9, 8, 10))
	if h.Length() == 10 {
		t.Log("ok")
	} else {
		t.Error("err")
	}

	h.Sort(func(a, b int) bool {
		return a > b
	})

	t.Log(h.Show())
}
