package numbers

import (
    "fmt"
    "sort"
    "strconv"
    "math"
)

const ConvertThreshold = 1000

var BinaryPrefixes = map[float64]string{
    1000000000000000: "%.2fP",
    1000000000000: "%.2fT",
    1000000000: "%.2fG",
    1000000: "%.2fM",
    1000: "%.1fk",
    0: "%.1f",
}

func getSizeFloatSlice() []float64 {
    floatSlice := make([]float64, 0, len(BinaryPrefixes))
    for num, _ := range BinaryPrefixes {
        floatSlice = append(floatSlice, num)
    }
    sort.Sort(sort.Reverse(sort.Float64Slice(floatSlice)))
    return floatSlice
}

func MetricSuffix(number float64) string {
    for size := range getSizeFloatSlice() {
        if size <= number {
            var value float64
            if number >= ConvertThreshold {
                value = number / size
            } else {
                value = number
            }
            
            return fmt.Sprintf(BinaryPrefixes[size], value)
        }
    }

    return strconv.FormatFloat(number, 'f', -1, 64)
}
