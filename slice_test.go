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
