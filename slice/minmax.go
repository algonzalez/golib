package slice

import "golang.org/x/exp/constraints"

// Max returns the largest item in slice.
func Max[T constraints.Ordered](slice []T) T {
	var max T
	for _, elem := range slice {
		if max < elem {
			max = elem
		}
	}
	return max
}

// Min returns the smallest item in slice.
// Min returns the smallest item in slice.
func Min[T constraints.Ordered](slice []T) T {
	var min T
	if len(slice) > 0 {
		min = slice[0]
		for _, elem := range slice {
			if min > elem {
				min = elem
			}
		}
	}
	return min
}
