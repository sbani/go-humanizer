package units

import (
	"fmt"
	"math"
	"sort"
)

// BinaryConvertThreshold is the number for a binary "step"
const BinaryConvertThreshold = 1024

// BinarySuffix returns the given number with a binary suffix.
func BinarySuffix(number float64) string {
	var str string
	if number < 0 {
		str = "-"
		number = math.Abs(number)
	}
	var binarySuffixPrefixes = map[float64]string{
		1125899906842624: "%.2f PB",
		1099511627776:    "%.2f TB",
		1073741824:       "%.2f GB",
		1048576:          "%.2f MB",
		1024:             "%.1f kB",
	}
	intSlice := make([]float64, 0, len(binarySuffixPrefixes))
	for num := range binarySuffixPrefixes {
		intSlice = append(intSlice, num)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(intSlice)))
	for _, size := range intSlice {
		if size <= number {
			var value float64
			if number >= BinaryConvertThreshold {
				value = number / float64(size)
			}
			return str + fmt.Sprintf(binarySuffixPrefixes[size], value)
		}
	}
	return fmt.Sprintf("%s%.0f bytes", str, number)
}
