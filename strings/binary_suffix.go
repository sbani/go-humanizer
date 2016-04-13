package strings

import (
	"fmt"
	"sort"
)

// BinaryConvertThreshold is the number for a binary "step"
const BinaryConvertThreshold = 1024

// BinarySuffixPrefixes are the floats with the specific binary suffix string fmt.
var BinarySuffixPrefixes = map[float64]string{
	1125899906842624: "%.2f PB",
	1099511627776:    "%.2f TB",
	1073741824:       "%.2f GB",
	1048576:          "%.2f MB",
	1024:             "%.1f kB",
	0:                "%.0f bytes",
}

// Get sorted slice from BinarySuffixPrefixes
func getIntSlice() []float64 {
	intSlice := make([]float64, 0, len(BinarySuffixPrefixes))
	for num := range BinarySuffixPrefixes {
		intSlice = append(intSlice, num)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(intSlice)))
	return intSlice
}

// BinarySuffix returns the given number with a binary suffix.
func BinarySuffix(number float64) string {
	for _, size := range getIntSlice() {
		if size <= number {
			var value float64
			if number >= BinaryConvertThreshold {
				value = number / float64(size)
			} else {
				value = number
			}

			return fmt.Sprintf(BinarySuffixPrefixes[size], value)
		}
	}

	return fmt.Sprintf("%f", number)
}
