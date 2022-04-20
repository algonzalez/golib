package golib

import "golang.org/x/exp/constraints"

// Max returns the larger of the two specified values.
func Max[T constraints.Ordered](x, y T) T {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of the two specified values.
func Min[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}
