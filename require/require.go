// Package require provides functions that to check conditions
// and panic when they are not met.
package require

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// IsZero enforces that a value is set to its zero value.
func IsZero[T comparable](x T) {
	var zero T
	if x != zero {
		panic(fmt.Sprintf("value must be the zero value for the %T type", x))
	}
}

// IsNotZero enforces that a value is NOT set to its zero value.
func IsNotZero[T comparable](x T) {
	var zero T
	if x == zero {
		panic(fmt.Sprintf("value must NOT be the zero value for the %T type", x))
	}
}

// IsTrue enforces that a condition is true.
func IsTrue(condition bool) {
	if !condition {
		panic("condition must be true; was false")
	}
}

// IsFalse enforces that a condition is false.
func IsFalse(condition bool) {
	if !condition {
		panic("condition must be false; was true")
	}
}

// IsInRange enforces that value is between low and high, exclusive of high.
//
// For example, 5 would meet the requirement with
// a low of 2 and a high of 10, but 10 would not.
func IsInRange[T constraints.Ordered](low, high, value T) {
	if !inRange(low, high, value) {
		panic(fmt.Sprintf("value must be between %v and %v, exclusive of %v; was %v", low, high, high, value))
	}
}

// IsIndexInRange enforces that index i is valid for the slice.
func IsIndexInRange[T any](i int, slice []T) {
	if slice == nil || !inRange(0, len(slice), i) {
		len := len(slice)
		panic(fmt.Sprintf("index must be between 0 and %d, exclusive of %d", len, len))
	}
}

// inRange determines if value is between low and high.
// It compares inclusive of low and exclusive of high, so [low..high)
func inRange[T constraints.Ordered](low, high, value T) bool {
	return value >= low && value < high
}
