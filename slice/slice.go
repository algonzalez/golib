package slice

import (
	"math"
)

// All returns true if all items in slice are matched by matches func;
func All[T comparable](slice []T, matches func(elem T) bool) bool {
	for _, elem := range slice {
		if !matches(elem) {
			return false
		}
	}
	return true
}

// Any returns true if at least one item in slice is matched by matches func.
func Any[T comparable](slice []T, matches func(elem T) bool) bool {
	for _, elem := range slice {
		if matches(elem) {
			return true
		}
	}
	return false
}

func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		panic("size must be an integer greater than zero")
	}

	itemCnt := len(slice)
	if itemCnt <= size {
		return [][]T{slice}
	}

	rowCnt := itemCnt / size
	lastRowSize := int(math.Mod(float64(itemCnt), float64(size)))
	if lastRowSize > 0 {
		rowCnt++
	}

	result := make([][]T, rowCnt)
	elemIndex := 0
	for i := range result {
		colSize := size
		if i == (rowCnt-1) && lastRowSize > 0 {
			colSize = lastRowSize
		}
		result[i] = make([]T, colSize)
		copy(result[i], slice[elemIndex:elemIndex+colSize])
		elemIndex += colSize
	}

	return result
}

// Compact removes all empty/zero values from slice.
func Compact[T comparable](slice []T) []T {
	result := []T{}
	var empty T
	for _, elem := range slice {
		if elem != empty {
			result = append(result, elem)
		}
	}
	return result
}

// Contains returns true if value exists in slice.
func Contains[T comparable](slice []T, value T) bool {
	for _, elem := range slice {
		if elem == value {
			return true
		}
	}
	return false
}

// ContainsAll returns true if all values exist in slice.
func ContainsAll[T comparable](slice []T, values ...T) bool {
	for _, value := range values {
		found := false
		for _, elem := range slice {
			if elem == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Contains returns true if any values exist in slice.
func ContainsAny[T comparable](slice []T, values ...T) bool {
	for _, value := range values {
		for _, elem := range slice {
			if elem == value {
				return true
			}
		}
	}
	return false
}

// Distinct returns a list of unique items in slice.
// It uses a map to determine uniqueness and the
// normalizeMapKey func to normalize the key for comparison.
func Distinct[T comparable](slice []T, normalizeMapKey func(T) T) []T {
	uniqueRef := make(map[T]bool, len(slice))
	var distinctSlice []T
	for _, elem := range slice {
		key := normalizeMapKey(elem)
		if _, found := uniqueRef[key]; !found {
			distinctSlice = append(distinctSlice, elem)
			uniqueRef[key] = true
		}
	}
	return distinctSlice
}

// Filter returns a list of items in slice that are matched by matches func.
func Filter[T comparable](slice []T, matches func(elem T) bool) []T {
	result := []T{}
	if matches == nil {
		return result
	}
	for _, elem := range slice {
		if matches(elem) {
			result = append(result, elem)
		}
	}
	return result
}

// First returns the first item in slice that is matched by matches func.
func First[T comparable](slice []T, matches func(elem T) bool) T {
	var result T
	if i := IndexOf(slice, matches); i >= 0 {
		result = slice[i]
	}
	return result
}

// ForEach calls func f with every item in slice
func ForEach[T any](slice []T, f func(elem T, index int)) {
	if f != nil {
		for i, elem := range slice {
			f(elem, i)
		}
	}
}

// IndexOf returns the index of the first item in slice
// that is matched by matches func.
func IndexOf[T comparable](slice []T, matches func(elem T) bool) int {
	for i, elem := range slice {
		if matches(elem) {
			return i
		}
	}
	return -1
}

func Insert[T any](slice []T, item T, i int) []T {
	switch {
	case i < 0:
		return []T{item}
	case i == 0:
		return append([]T{item}, slice[i:]...)
	case i >= len(slice):
		return append(slice, item)
	}
	return append(slice[:i], append([]T{item}, slice[i:]...)...)
}

// Last returns the last item in slice that is matched by matches func.
func Last[T comparable](slice []T, matches func(elem T) bool) T {
	var result T
	if i := LastIndexOf(slice, matches); i >= 0 {
		result = slice[i]
	}
	return result
}

// LastIndexOf returns the index of the last item in slice
// that is matched by matches func.
func LastIndexOf[T comparable](slice []T, matches func(elem T) bool) int {
	for i := len(slice); i >= 0; i-- {
		if matches(slice[i]) {
			return i
		}
	}
	return -1
}

func Map[T comparable](slice []T, mapper func(elem T) T) []T {
	if mapper == nil {
		return []T{}
	}
	result := make([]T, len(slice))
	for i, elem := range slice {
		result[i] = mapper(elem)
	}
	return result
}

func Prepend[T any](slice []T, item T) []T {
	return append([]T{item}, slice[0:]...)
}

func Reduce[T comparable, TResult any](slice []T, seed TResult, reducer func(acc TResult, elem T, index int) TResult) TResult {
	result := seed
	if reducer == nil {
		return result
	}
	for i, elem := range slice {
		result = reducer(result, elem, i)
	}
	return result
}

func Reverse[T any](slice []T) []T {
	inputLen := len(slice)
	if inputLen == 0 {
		return slice
	}
	result := make([]T, inputLen)
	j := 0
	for i := inputLen - 1; i >= 0; i-- {
		result[j] = slice[i]
		j++
	}
	return result
}
