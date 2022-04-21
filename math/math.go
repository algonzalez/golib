package math

import (
	"fmt"
	"math"
	"strconv"
)

// PiString returns a string representation of math.Pi (a.k.a. π),
// where digits specifies the number of decimal places to include.
// The digits value must be in the range of 0..48 inclusive of 48.
func PiString(digits int) string {
	if digits < 0 {
		digits = 0
	} else if digits > 48 {
		digits = 48
	}
	return fmt.Sprintf("%."+strconv.Itoa(digits)+"f", math.Pi)
}
