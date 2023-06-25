package slices

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args[T any] struct {
		a [][]T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test2",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		a []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeUnique(t *testing.T) {
	type args[T comparable] struct {
		a [][]T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test2",
			args: args[int]{
				a: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeUnique(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args[T any] struct {
		a []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			},
			want: []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args[T any] struct {
		a    []T
		less func(i, j int) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				less: func(i, j int) bool {
					return false
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				less: func(i, j int) bool {
					return true
				},
			},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.a, tt.args.less); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortBy(t *testing.T) {
	type args[T any] struct {
		a   []T
		key func(front, back T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type Student struct {
		Name string
		Age  int
	}

	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{
						Name: "a",
						Age:  1,
					}, {
						Name: "b",
						Age:  2,
					},
				},
				key: func(front, back Student) bool {
					return front.Age < back.Age
				},
			},
			want: []Student{
				{
					Name: "a",
					Age:  1,
				}, {
					Name: "b",
					Age:  2,
				},
			},
		},
		{
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				key: func(front, back Student) bool {
					return front.Age > back.Age
				},
			},
			want: []Student{
				{Name: "b", Age: 2},
				{Name: "a", Age: 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortBy(tt.args.a, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 1,
			},
			want: 0,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 9,
			},
			want: 8,
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Index(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexBy(t *testing.T) {
	type args[T any] struct {
		a   []T
		key func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want int
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				key: func(v Student) bool {
					return v.Age == 1
				},
			},
			want: 0,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				key: func(v Student) bool {
					return v.Age == 3
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IndexBy(tt.args.a, tt.args.key); got != tt.want {
				t.Errorf("IndexBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndex(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 1,
			},
			want: 0,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
				v: 1,
			},
			want: 9,
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndex(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("LastIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastIndexBy(t *testing.T) {
	type args[T any] struct {
		a   []T
		key func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want int
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
					{Name: "c", Age: 1},
				},
				key: func(v Student) bool {
					return v.Age == 1
				},
			},
			want: 2,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
					{Name: "c", Age: 1},
				},
				key: func(v Student) bool {
					return v.Age == 3
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastIndexBy(tt.args.a, tt.args.key); got != tt.want {
				t.Errorf("LastIndexBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 1,
			},
			want: true,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsBy(t *testing.T) {
	type args[T any] struct {
		a   []T
		key func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
					{Name: "c", Age: 1},
				},
				key: func(v Student) bool {
					return v.Age == 1
				},
			},
			want: true,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
					{Name: "c", Age: 1},
				},
				key: func(v Student) bool {
					return v.Age == 3
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsBy(tt.args.a, tt.args.key); got != tt.want {
				t.Errorf("ContainsBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: []int{1, 2, 3},
			},
			want: true,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: []int{1, 2, 10},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.a, tt.args.v...); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAllBy(t *testing.T) {
	type args[T any] struct {
		a    []T
		keys []func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}

	type Student struct {
		Name string
		Age  int
	}

	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				keys: []func(v Student) bool{
					func(v Student) bool {
						return v.Age == 1
					},
					func(v Student) bool {
						return v.Age == 2
					},
				},
			},
			want: true,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				keys: []func(v Student) bool{
					func(v Student) bool {
						return v.Age == 1
					},
					func(v Student) bool {
						return v.Age == 3
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAllBy(tt.args.a, tt.args.keys...); got != tt.want {
				t.Errorf("ContainsAllBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAny(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: []int{1, 2, 3},
			},
			want: true,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: []int{10, 11, 12},
			},
			want: false,
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: []int{10, 11, 12, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAny(tt.args.a, tt.args.v...); got != tt.want {
				t.Errorf("ContainsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAnyBy(t *testing.T) {
	type args[T any] struct {
		a    []T
		keys []func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				keys: []func(v Student) bool{
					func(v Student) bool {
						return v.Age == 1
					},
					func(v Student) bool {
						return v.Age == 3
					},
				},
			},
			want: true,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				keys: []func(v Student) bool{
					func(v Student) bool {
						return v.Age == 3
					},
					func(v Student) bool {
						return v.Age == 4
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAnyBy(tt.args.a, tt.args.keys...); got != tt.want {
				t.Errorf("ContainsAnyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 1,
			},
			want: 1,
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				v: 10,
			},
			want: 0,
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
				v: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountBy(t *testing.T) {
	type args[T any] struct {
		a   []T
		key func(v T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want int
	}

	type Student struct {
		Name string
		Age  int
	}

	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				key: func(v Student) bool {
					return v.Age == 1
				},
			},
			want: 1,
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				key: func(v Student) bool {
					return v.Age == 3
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBy(tt.args.a, tt.args.key); got != tt.want {
				t.Errorf("CountBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		a []T
		f func(T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, float64]{
		{
			name: "test1",
			args: args[int, float64]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f: func(v int) float64 {
					return float64(v)
				},
			},
			want: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
					{Name: "c", Age: 3},
				},
				f: func(v Student) bool {
					return v.Age > 1
				},
			},
			want: []Student{
				{Name: "b", Age: 2},
				{Name: "c", Age: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantRes T
	}

	type Student struct {
		Name string
		Age  int
	}
	tests := []testCase[Student]{
		{
			name: "test1",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				f: func(v Student) bool {
					return v.Age == 1
				},
			},
			wantRes: Student{Name: "a", Age: 1},
		}, {
			name: "test2",
			args: args[Student]{
				a: []Student{
					{Name: "a", Age: 1},
					{Name: "b", Age: 2},
				},
				f: func(v Student) bool {
					return v.Age == 3
				},
			},
			wantRes: Student{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Find(tt.args.a, tt.args.f); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Find() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	type args[T any] struct {
		a    []T
		size int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				size: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.a, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args[T any] struct {
		v []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				v: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "test2",
			args: args[int]{
				v: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args[T any] struct {
		a []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args[T any] struct {
		a [][]T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			name: "test2",
			args: args[int]{
				a: [][]int{
					{},
					{1, 2, 3},
					{},
					{4, 5, 6},
				},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Append(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	type args[T any] struct {
		a []T
		b [][]T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				b: [][]int{
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{4, 5, 6, 7, 8, 9, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prepend(tt.args.a, tt.args.b...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prepend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	type args[T any] struct {
		a []T
		i int
		v []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 1,
				v: []int{4, 5, 6},
			},
			want: []int{1, 4, 5, 6, 2, 3},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 0,
				v: []int{4, 5, 6},
			},
			want: []int{4, 5, 6, 1, 2, 3},
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 3,
				v: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.a, tt.args.i, tt.args.v...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertAll(t *testing.T) {
	type args[T any] struct {
		a []T
		i int
		v []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 1,
				v: []int{4, 5, 6},
			},
			want: []int{1, 4, 5, 6, 2, 3},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 0,
				v: []int{4, 5, 6},
			},
			want: []int{4, 5, 6, 1, 2, 3},
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 3,
				v: []int{4, 5, 6},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAll(tt.args.a, tt.args.i, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args[T any] struct {
		a []T
		i int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 1,
			},
			want: []int{1, 3},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 0,
			},
			want: []int{2, 3},
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 2,
			},
			want: []int{1, 2},
		}, {
			name: "test4",
			args: args[int]{
				a: []int{1, 2, 3},
				i: 3,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.a, tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAll(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 2,
			},
			want: []int{1, 3},
		}, {
			name: "test2",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 1,
			},
			want: []int{2, 3},
		}, {
			name: "test3",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 3,
			},
			want: []int{1, 2},
		}, {
			name: "test4",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 4,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAll(tt.args.a, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIf(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type T1 struct {
		a int
	}
	tests := []testCase[T1]{
		{
			name: "test1",
			args: args[T1]{
				a: []T1{{1}, {2}, {3}},
				f: func(t T1) bool {
					return t.a == 2
				},
			},
			want: []T1{{1}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIf(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFirst(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 2, 3},
				v: 2,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirst(tt.args.a, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveLast(t *testing.T) {
	type args[T comparable] struct {
		a []T
		v T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3, 2},
				v: 2,
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLast(tt.args.a, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveFirstIf(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}

	type T1 struct {
		a int
	}
	tests := []testCase[T1]{
		{
			name: "test1",
			args: args[T1]{
				a: []T1{{1}, {2}, {2}, {3}},
				f: func(t T1) bool {
					return t.a == 2
				},
			},
			want: []T1{{1}, {2}, {3}},
		}, {
			name: "test2",
			args: args[T1]{
				a: []T1{{1}, {2}, {3}},
				f: func(t T1) bool {
					return t.a == 1
				},
			},
			want: []T1{{2}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirstIf(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirstIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveLastIf(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	type T1 struct {
		a int
	}
	tests := []testCase[T1]{
		{
			name: "test1",
			args: args[T1]{
				a: []T1{{1}, {2}, {3}, {2}},
				f: func(t T1) bool {
					return t.a == 2
				},
			},
			want: []T1{{1}, {2}, {3}},
		}, {
			name: "test2",
			args: args[T1]{
				a: []T1{{1}, {2}, {3}},
				f: func(t T1) bool {
					return t.a == 1
				},
			},
			want: []T1{{2}, {3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveLastIf(tt.args.a, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveLastIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRange(t *testing.T) {
	type args[T any] struct {
		a    []T
		from int
		to   int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3, 4, 5},
				from: 1,
				to:   3,
			},
			want: []int{1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRange(tt.args.a, tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(int, T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				f: func(i int, t int) {

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Range(tt.args.a, tt.args.f)
		})
	}
}

func TestRangeReverse(t *testing.T) {
	type args[T any] struct {
		a []T
		f func(int, T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				f: func(i int, t int) {

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeReverse(tt.args.a, tt.args.f)
		})
	}
}

func TestRangeRange(t *testing.T) {
	type args[T any] struct {
		a    []T
		from int
		to   int
		f    func(int, T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3},
				from: 1,
				to:   2,
				f: func(i int, t int) {

				},
			},
		}, {
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3},
				from: 2,
				to:   2,
				f: func(i int, t int) {

				},
			},
		}, {
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3},
				from: 3,
				to:   2,
				f: func(i int, t int) {

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeRange(tt.args.a, tt.args.from, tt.args.to, tt.args.f)
		})
	}
}

func TestRangeRangeReverse(t *testing.T) {
	type args[T any] struct {
		a    []T
		from int
		to   int
		f    func(int, T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a:    []int{1, 2, 3},
				from: 1,
				to:   2,
				f: func(i int, t int) {

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RangeRangeReverse(tt.args.a, tt.args.from, tt.args.to, tt.args.f)
		})
	}
}

func TestReduce(t *testing.T) {
	type args[T any, U any] struct {
		a       []T
		f       func(U, T) U
		initial U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, float64]{
		{
			name: "test1",
			args: args[int, float64]{
				a:       []int{1, 2, 3},
				f:       func(u float64, t int) float64 { return u + float64(t) },
				initial: 0,
			},
			want: 6,
		}, {
			name: "test1",
			args: args[int, float64]{
				a:       []int{1, 2, 3},
				f:       func(u float64, t int) float64 { return u - float64(t) },
				initial: 0,
			},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.a, tt.args.f, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceReverse(t *testing.T) {
	type args[T any, U any] struct {
		a       []T
		f       func(U, T) U
		initial U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, float64]{
		{
			name: "test1",
			args: args[int, float64]{
				a:       []int{1, 2, 3},
				f:       func(u float64, t int) float64 { return u + float64(t) },
				initial: 0,
			},
			want: 6,
		}, {
			name: "test1",
			args: args[int, float64]{
				a:       []int{1, 2, 3},
				f:       func(u float64, t int) float64 { return u - float64(t) },
				initial: 0,
			},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceReverse(tt.args.a, tt.args.f, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReduceReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushFront(t *testing.T) {
	type args[T any] struct {
		a []T
		v T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 4,
			},
			want: []int{4, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PushFront(tt.args.a, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPushBack(t *testing.T) {
	type args[T any] struct {
		a []T
		v T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				v: 4,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PushBack(tt.args.a, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PushBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopFront(t *testing.T) {
	type args[T any] struct {
		a []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantArr []T
		wantR   T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
			},
			wantArr: []int{2, 3},
			wantR:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArr, gotR := PopFront(tt.args.a)
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("PopFront() gotArr = %v, want %v", gotArr, tt.wantArr)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("PopFront() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestPopBack(t *testing.T) {
	type args[T any] struct {
		a []T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantArr []T
		wantR   T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
			},
			wantArr: []int{1, 2},
			wantR:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArr, gotR := PopBack(tt.args.a)
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("PopBack() gotArr = %v, want %v", gotArr, tt.wantArr)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("PopBack() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestPopFrontN(t *testing.T) {
	type args[T any] struct {
		a []T
		n int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantArr []T
		wantR   []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				n: 2,
			},
			wantArr: []int{3},
			wantR:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArr, gotR := PopFrontN(tt.args.a, tt.args.n)
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("PopFrontN() gotArr = %v, want %v", gotArr, tt.wantArr)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("PopFrontN() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestPopBackN(t *testing.T) {
	type args[T any] struct {
		a []T
		n int
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantArr []T
		wantR   []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a: []int{1, 2, 3},
				n: 2,
			},
			wantArr: []int{1},
			wantR:   []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArr, gotR := PopBackN(tt.args.a, tt.args.n)
			if !reflect.DeepEqual(gotArr, tt.wantArr) {
				t.Errorf("PopBackN() gotArr = %v, want %v", gotArr, tt.wantArr)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("PopBackN() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	type args[T any] struct {
		a     []T
		start int
		end   []int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "test1",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 1,
				end:   []int{3},
			},
			want: []int{2, 3, 4},
		}, {
			name: "test2",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 1,
				end:   []int{2, 4},
			},
			want: []int{2, 3},
		}, {
			name: "test3",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: -1,
				end:   []int{3, 4},
			},
			want: []int{5},
		}, {
			name: "test4",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: -1,
				end:   []int{-3},
			},
			want: []int{5, 4, 3},
		}, {
			name: "test5",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 1,
				end:   []int{-3},
			},
			want: []int{2, 1},
		}, {
			name: "test6",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 1,
				end:   []int{3, 4, 5},
			},
			want: []int{2, 3, 4},
		}, {
			name: "test7",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: -2,
				end:   []int{3, 4, 5},
			},
			want: []int{4, 5},
		}, {
			name: "test8",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 2,
				end:   []int{-3},
			},
			want: []int{3, 2, 1},
		}, {
			name: "test9",
			args: args[int]{
				a:     []int{1, 2, 3, 4, 5},
				start: 2,
				end:   []int{3, 4, 5},
			},
			want: []int{3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.a, tt.args.start, tt.args.end...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
