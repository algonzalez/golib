package require

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AreEqual enforces that two values are equal to each other.
func AreEqual[T comparable](x, y T) {
	if x != y {
		panic(fmt.Sprintf("values must be equal\n\tfirst: %v\n\tsecond: %v\n", x, y))
	}
}

// AreNotEqual enforces that two value are NOT equal to each other.
func AreNotEqual[T comparable](x, y T) {
	if x == y {
		panic(fmt.Sprintf("values must NOT be equal\n\tfirst: %v\n\tsecond: %v\n", x, y))
	}
}

// IsLessThan enforces that value < target.
func IsLessThan[T constraints.Ordered](target, value T) {
	if value >= target {
		panic(fmt.Sprintf("value must be less than target\n\tvalue: %v\n\ttarget: %v", value, target))
	}
}

// IsLessThanOrEquals enforces that value <= target.
func IsLessThanOrEquals[T constraints.Ordered](target, value T) {
	if value > target {
		panic(fmt.Sprintf("value must be less than or equal to target\n\tvalue: %v\n\ttarget: %v", value, target))
	}
}

// IsGreaterThan enforces that value > target.
func IsGreaterThan[T constraints.Ordered](target, value T) {
	if value <= target {
		panic(fmt.Sprintf("value must be greater than target\n\tvalue: %v\n\ttarget: %v", value, target))
	}
}

// IsGreaterThanOrEquals enforces that value >= target.
func IsGreaterThanOrEquals[T constraints.Ordered](target, value T) {
	if value < target {
		panic(fmt.Sprintf("value must be greater than or equal to target\n\tvalue: %v\n\ttarget: %v", value, target))
	}
}
