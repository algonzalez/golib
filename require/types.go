package require

import (
	"fmt"
	"reflect"
)

// IsOfTYpe enforces that value is of type t.
func IsOfType[T any](value T, t reflect.Type) {
	vt := reflect.TypeOf(value)
	if t != vt {
		panic(fmt.Sprintf("value must be of type %v; was %v", t, vt))
	}
}

// IsOfTYpe enforces that value is of type bool.
func IsOfTypeBool(value interface{}) {
	IsOfType(value, reflect.TypeOf(true))
}

// IsOfTYpe enforces that value is of type int.
func IsOfTypeInt(value interface{}) {
	IsOfType(value, reflect.TypeOf(0))
}

// IsOfTYpe enforces that value is of type float32.
func IsOfTypeFloat32(value interface{}) {
	IsOfType(value, reflect.TypeOf(float32(0.0)))
}

// IsOfTYpe enforces that value is of type float64.
func IsOfTypeFloat64(value interface{}) {
	IsOfType(value, reflect.TypeOf(float64(0.0)))
}

// IsOfTYpe enforces that value is of type string.
func IsOfTypeString(value interface{}) {
	IsOfType(value, reflect.TypeOf(""))
}
