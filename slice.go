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

// NewList returns a new slice containing the specified values.
//
//	NewList 返回一个包含指定值的新切片
func NewList[T any](v ...T) []T {
	return v
}

// Copy returns a new slice containing the elements of the specified slice.
//
//	Copy 返回一个包含指定切片元素的新切片
func Copy[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

// Append returns a new slice containing the elements of the specified slices.
//
//	Append 返回一个包含指定切片元素的新切片
func Append[T any](a ...[]T) []T {
	var n int
	for _, x := range a {
		n += len(x)
	}
	b := make([]T, 0, n)
	for _, x := range a {
		b = append(b, x...)
	}
	return b
}

// Prepend returns a new slice containing the elements of the specified slices.
//
//	Prepend 返回一个包含指定切片元素的新切片, 与Append插入顺序相反
func Prepend[T any](a []T, b ...[]T) []T {
	var n int
	for _, x := range b {
		n += len(x)
	}
	c := make([]T, 0, len(a)+n)
	for _, x := range b {
		c = append(c, x...)
	}
	c = append(c, a...)
	return c
}

// Insert returns a new slice containing the elements of the specified slice with the specified values inserted at the specified index.
//
//	Insert 返回一个包含指定切片元素的新切片，该切片在指定索引处插入指定值
func Insert[T any](a []T, index int, v ...T) []T {
	if index < 0 || index > len(a) {
		panic("index out of range")
	}

	return append(a[:index], append(v, a[index:]...)...)
}

// InsertAll returns a new slice containing the elements of the specified slice with the specified values inserted at the specified index.
//
//	InsertAll 返回一个包含指定切片元素的新切片，该切片在指定索引处插入指定切片的所有元素
func InsertAll[T any](a []T, index int, v []T) []T {
	return Insert(a, index, v...)
}

// Remove returns a new slice containing the elements of the specified slice with the element at the specified index removed.
//
//	Remove 返回一个包含指定切片元素的新切片，该切片删除指定索引处的元素
func Remove[T any](a []T, index int) []T {
	if index < 0 || index >= len(a) {
		panic("index out of range")
	}

	return append(a[:index], a[index+1:]...)
}

// RemoveAll returns a new slice containing the elements of the specified slice with all elements equal to the specified value removed.
//
//	RemoveAll 返回一个包含指定切片元素的新切片，该切片删除所有等于指定值的元素
func RemoveAll[T comparable](a []T, v T) []T {
	b := make([]T, 0, len(a))
	for _, x := range a {
		if x != v {
			b = append(b, x)
		}
	}
	return b
}

// RemoveIf returns a new slice containing the elements of the specified slice with all elements for which the specified function returns true removed.
//
//	RemoveIf 返回一个包含指定切片元素的新切片，该切片删除指定函数返回true的所有元素
func RemoveIf[T any](a []T, f func(T) bool) []T {
	b := make([]T, 0, len(a))
	for _, x := range a {
		if !f(x) {
			b = append(b, x)
		}
	}
	return b
}

// RemoveFirst returns a new slice containing the elements of the specified slice with the first element equal to the specified value removed.
//
//	RemoveFirst 返回一个包含指定切片元素的新切片，该切片删除第一个等于指定值的元素
func RemoveFirst[T comparable](a []T, v T) []T {
	for i, x := range a {
		if x == v {
			return Remove(a, i)
		}
	}
	return a
}

// RemoveLast returns a new slice containing the elements of the specified slice with the last element equal to the specified value removed.
//
//	RemoveLast 返回一个包含指定切片元素的新切片，该切片删除最后一个等于指定值的元素
func RemoveLast[T comparable](a []T, v T) []T {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == v {
			return Remove(a, i)
		}
	}
	return a
}

// RemoveFirstIf returns a new slice containing the elements of the specified slice with the first element for which the specified function returns true removed.
//
//	RemoveFirstIf 返回一个包含指定切片元素的新切片，该切片删除第一个指定函数返回true的元素
func RemoveFirstIf[T any](a []T, f func(T) bool) []T {
	for i, x := range a {
		if f(x) {
			return Remove(a, i)
		}
	}
	return a
}

// RemoveLastIf returns a new slice containing the elements of the specified slice with the last element for which the specified function returns true removed.
//
//	RemoveLastIf 返回一个包含指定切片元素的新切片，该切片删除最后一个指定函数返回true的元素
func RemoveLastIf[T any](a []T, f func(T) bool) []T {
	for i := len(a) - 1; i >= 0; i-- {
		if f(a[i]) {
			return Remove(a, i)
		}
	}
	return a
}

// RemoveRange returns a new slice containing the elements of the specified slice with the elements in the specified range removed.
//
//	RemoveRange 返回一个包含指定切片元素的新切片，该切片删除指定范围内的元素
func RemoveRange[T any](a []T, from, to int) []T {
	if from < 0 || from > len(a) {
		panic("from index out of range")
	}
	if to < 0 || to > len(a) {
		panic("to index out of range")
	}
	if from > to {
		panic("from index greater than to index")
	}

	return append(a[:from], a[to:]...)
}

// Range iterates over all elements of the specified slice.
//
//	Range 遍历指定切片所有元素, 解决for循环中共用变量问题
func Range[T any](a []T, f func(int, T)) {
	for i, x := range a {
		temp := x
		f(i, temp)
	}
}

// RangeReverse iterates over all elements of the specified slice in reverse order.
//
//	RangeReverse 反向遍历指定切片所有元素, 解决for循环中共用变量问题
func RangeReverse[T any](a []T, f func(int, T)) {
	for i := len(a) - 1; i >= 0; i-- {
		temp := a[i]
		f(i, temp)
	}
}

// RangeRange iterates over all elements of the specified slice in the specified range.
//
//	RangeRange 遍历指定切片指定范围内的所有元素, 解决for循环中共用变量问题
func RangeRange[T any](a []T, from, to int, f func(int, T)) {
	if from < 0 || from > len(a) {
		panic("from index out of range")
	}
	if to < 0 || to > len(a) {
		panic("to index out of range")
	}
	if from > to {
		panic("from index greater than to index")
	}

	for i := from; i < to; i++ {
		temp := a[i]
		f(i, temp)
	}
}

// RangeRangeReverse iterates over all elements of the specified slice in the specified range in reverse order.
//
//	RangeRangeReverse 反向遍历指定切片指定范围内的所有元素, 解决for循环中共用变量问题
func RangeRangeReverse[T any](a []T, from, to int, f func(int, T)) {
	if from < 0 || from > len(a) {
		panic("from index out of range")
	}
	if to < 0 || to > len(a) {
		panic("to index out of range")
	}
	if from > to {
		panic("from index greater than to index")
	}

	for i := to - 1; i >= from; i-- {
		temp := a[i]
		f(i, temp)
	}
}

// Reduce reduces the specified slice to a single value by iteratively combining each element of the slice with the running result using the specified function.
//
//	Reduce 通过使用指定函数将切片的每个元素与运行结果迭代地组合来将指定切片减少为单个值
func Reduce[T, U any](a []T, f func(U, T) U, initial U) U {
	result := initial
	for _, x := range a {
		result = f(result, x)
	}
	return result
}

// ReduceReverse reduces the specified slice to a single value by iteratively combining each element of the slice in reverse order with the running result using the specified function.
//
//	ReduceReverse 反向通过使用指定函数将切片的每个元素与运行结果迭代地组合来将指定切片减少为单个值
func ReduceReverse[T, U any](a []T, f func(U, T) U, initial U) U {
	result := initial
	for i := len(a) - 1; i >= 0; i-- {
		result = f(result, a[i])
	}
	return result
}
