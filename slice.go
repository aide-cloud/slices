package slices

import "sort"

// Merge merges the specified slices into a new slice.
//
//	Merge 合并指定的切片到一个新的切片中
func Merge[T any](a ...[]T) []T {
	var b []T
	for _, x := range a {
		b = append(b, x...)
	}
	return b
}

// Unique returns a new slice containing the elements of the specified slice with duplicates removed.
//
//	Unique 返回一个新的切片，该切片包含指定切片中的元素，去除重复元素
func Unique[T comparable](a []T) []T {
	b := make([]T, 0, len(a)/2)
	tmp := make(map[T]struct{}, len(a)/2)
	for _, x := range a {
		if _, ok := tmp[x]; !ok {
			tmp[x] = struct{}{}
			b = append(b, x)
		}
	}

	return b
}

// MergeUnique merges the specified slices into a new slice with duplicates removed.
//
//	MergeUnique 合并指定的切片到一个新的切片中，去除重复元素
func MergeUnique[T comparable](a ...[]T) []T {
	return Unique(Merge(a...))
}

// Reverse returns a new slice containing the elements of the specified slice in reverse order.
//
//	Reverse 返回一个新的切片，该切片包含指定切片中的元素，元素顺序反转
func Reverse[T any](a []T) []T {
	b := make([]T, len(a))
	// 首尾元素交换
	for i, j := 0, len(a)-1; i < len(a)/2+1; i, j = i+1, j-1 {
		b[i], b[j] = a[j], a[i]
	}
	return b
}

// Sort returns a new slice containing the elements of the specified slice in sorted order.
//
//	Sort 返回一个新的切片，该切片包含指定切片中的元素，元素按照指定的比较函数排序
func Sort[T any](a []T, less func(i, j int) bool) []T {
	b := make([]T, len(a))
	copy(b, a)
	sort.Slice(b, less)
	return b
}

// SortBy returns a new slice containing the elements of the specified slice in sorted order.
//
//	SortBy 返回一个新的切片，该切片包含指定切片中的元素，元素按照指定的比较函数排序
func SortBy[T any](a []T, key func(front, back T) bool) []T {
	return Sort(a, func(i, j int) bool {
		return key(a[i], a[j])
	})
}

// Index returns the index of the first instance of the specified value in the specified slice, or -1 if not found.
//
//	Index 返回指定切片中第一个指定值的索引，如果没有找到则返回-1
func Index[T comparable](a []T, v T) int {
	for i, x := range a {
		if x == v {
			return i
		}
	}
	return -1
}

// IndexBy returns the index of the first instance of the specified value in the specified slice, or -1 if not found.
//
//	IndexBy 返回指定切片中第一个指定值的索引，如果没有找到则返回-1
func IndexBy[T any](a []T, key func(v T) bool) int {
	for i, x := range a {
		if key(x) {
			return i
		}
	}
	return -1
}

// LastIndex returns the index of the last instance of the specified value in the specified slice, or -1 if not found.
//
//	LastIndex 返回指定切片中最后一个指定值的索引，如果没有找到则返回-1
func LastIndex[T comparable](a []T, v T) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == v {
			return i
		}
	}
	return -1
}

// LastIndexBy returns the index of the last instance of the specified value in the specified slice, or -1 if not found.
//
//	LastIndexBy 返回指定切片中最后一个指定值的索引，如果没有找到则返回-1
func LastIndexBy[T any](a []T, key func(v T) bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if key(a[i]) {
			return i
		}
	}
	return -1
}

// Contains returns true if the specified slice contains the specified value.
//
//	Contains 如果指定切片包含指定值，则返回true
func Contains[T comparable](a []T, v T) bool {
	return Index(a, v) >= 0
}

// ContainsBy returns true if the specified slice contains the specified value.
//
//	ContainsBy 如果指定切片包含指定值，则返回true
func ContainsBy[T any](a []T, key func(v T) bool) bool {
	return IndexBy(a, key) >= 0
}

// ContainsAll returns true if the specified slice contains all the specified values.
//
//	ContainsAll 如果指定切片包含所有指定值，则返回true
func ContainsAll[T comparable](a []T, v ...T) bool {
	for _, x := range v {
		if !Contains(a, x) {
			return false
		}
	}
	return true
}

// ContainsAllBy returns true if the specified slice contains all the specified values.
//
//	ContainsAllBy 如果指定切片包含所有指定值，则返回true
func ContainsAllBy[T any](a []T, keys ...func(v T) bool) bool {
	for _, key := range keys {
		if !ContainsBy(a, key) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the specified slice contains any of the specified values.
//
//	ContainsAny 如果指定切片包含任意一个指定值，则返回true
func ContainsAny[T comparable](a []T, v ...T) bool {
	for _, x := range v {
		if Contains(a, x) {
			return true
		}
	}
	return false
}

// ContainsAnyBy returns true if the specified slice contains any of the specified values.
//
//	ContainsAnyBy 如果指定切片包含任意一个指定值，则返回true
func ContainsAnyBy[T any](a []T, keys ...func(v T) bool) bool {
	for _, key := range keys {
		if ContainsBy(a, key) {
			return true
		}
	}
	return false
}

// Count returns the number of elements in the specified slice that equal the specified value.
//
//	Count 返回指定切片中等于指定值的元素数量
func Count[T comparable](a []T, v T) int {
	c := 0
	for _, x := range a {
		if x == v {
			c++
		}
	}
	return c
}

// CountBy returns the number of elements in the specified slice that equal the specified value.
//
//	CountBy 返回指定切片中等于指定值的元素数量
func CountBy[T any](a []T, key func(v T) bool) int {
	c := 0
	for _, x := range a {
		if key(x) {
			c++
		}
	}
	return c
}

// Map returns a new slice containing the results of applying the specified function to each element of the specified slice.
//
//	Map 返回一个新的切片，该切片包含将指定函数应用于指定切片的每个元素的结果
func Map[T, U any](a []T, f func(T) U) []U {
	b := make([]U, len(a))
	for i, x := range a {
		b[i] = f(x)
	}
	return b
}

// Filter returns a new slice containing the elements of the specified slice for which the specified function returns true.
//
//	Filter 返回一个新的切片，该切片包含指定切片中指定函数返回true的元素
func Filter[T any](a []T, f func(T) bool) []T {
	b := make([]T, 0, len(a))
	for _, x := range a {
		if f(x) {
			b = append(b, x)
		}
	}
	return b
}

// Find returns the first element of the specified slice for which the specified function returns true, or nil if not found.
//
//	Find 返回指定切片中指定函数返回true的第一个元素，如果没有找到则返回nil
func Find[T any](a []T, f func(T) bool) (res T) {
	for _, x := range a {
		if f(x) {
			return x
		}
	}
	return
}

// Split splits the specified slice into a slice of slices, each containing size elements, with the final slice containing the remainder.
//
//	Split 将指定切片分割成一个包含size个元素的切片的切片，最后一个切片包含剩余的元素
func Split[T any](a []T, size int) [][]T {
	if size <= 0 {
		panic("size must be positive")
	}
	b := make([][]T, 0, (len(a)+size-1)/size)
	for len(a) > size {
		b = append(b, a[:size:size])
		a = a[size:]
	}
	return append(b, a)
}
