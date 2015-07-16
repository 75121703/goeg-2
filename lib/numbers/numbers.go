package numbers

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

// Uint8FromInt accepts an int argument and returns a uint8 and nil if the int is in range,
// or 0 and an error otherwise.
func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}

	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}

// IntFromFloat64 accepts a float64 and returns a int and nil if the float64 is in range,
// or 0 and an error otherwise. Rather than simply returning the whole part (i.e., truncating),
// this function perform a very simple rounding if the fractional part is >= 0.5.
func IntFromFloat64(x float64) (int, error) {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)

		if fraction >= 0.5 {
			whole++
		}

		return int(whole), nil
	}

	return 0, fmt.Errorf("%g is out of the int32 range", x)
}

// EqualFloat compares two float64s to the given accuracy - or to the gratest accuracy
// the machine can achieve if a negative number (e.g., -1) is passed as the limit.
func EqualFloat(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64
	}

	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

// EqualComplex compares two complex128 numbers
func EqualComplex(x, y complex128) bool {
	return EqualFloat(real(x), real(y), -1) && EqualFloat(imag(x), imag(y), -1)
}

// IntForBool converts bool to int
func IntForBool(b bool) int {
	if b {
		return 1
	}

	return 0
}

// Pad pads a number with the given width and rune
func Pad(number, width int, pad rune) string {
	s := fmt.Sprint(number)
	gap := width - utf8.RuneCountInString(s)

	if gap > 0 {
		return strings.Repeat(string(pad), gap) + s
	}

	return s
}

// Humanize returns a string representation of the number it is given with grouping separators (for
// languages that use simple three-digit groups) and padding
func Humanize(amount float64, width, decimals int, pad, separator rune) string {
	dollars, cents := math.Modf(amount)
	whole := fmt.Sprintf("%+.0f", dollars)[1:] // strip "±"
	fraction := ""
	if decimals > 0 {
		fraction = fmt.Sprintf("%+.*f", decimals, cents)[2:] // strip "±0"
	}

	sep := string(separator)
	for i := len(whole) - 3; i > 0; i -= 3 {
		whole = whole[:i] + sep + whole[i:]
	}
	if amount < 0.0 {
		whole = "-" + whole
	}

	number := whole + fraction
	gap := width - utf8.RuneCountInString(number)
	if gap > 0 {
		return strings.Repeat(string(pad), gap) + number
	}
	return number
}
