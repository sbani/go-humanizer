package strings

import (
    "sort"
    "fmt"
)

const BinaryConvertThreshold = 1024

var BinarySuffixPrefixes = map[float64]string{
    1125899906842624: "%.2f PB",
    1099511627776: "%.2f TB",
    1073741824: "%.2f GB",
    1048576: "%.2f MB",
    1024: "%.1f kB",
    0: "%.0f bytes",
}

// Get sorted slice from BinarySuffixPrefixes
func getIntSlice() []float64 {
    intSlice := make([]float64, 0, len(BinarySuffixPrefixes))
    for num, _ := range BinarySuffixPrefixes {
        intSlice = append(intSlice, num)
    }
    sort.Sort(sort.Reverse(sort.Float64Slice(intSlice)))
    return intSlice
}

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
