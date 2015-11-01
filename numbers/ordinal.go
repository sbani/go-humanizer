package numbers

import (
    "math"
    "strconv"
)

// Get the ordinal only for the specific number.
func Ordinal(number int) string {
    absNumber := int(math.Abs(float64(number)))
    
    i := absNumber % 100
    if i == 11 || i == 12 || i == 13 {
        return "th"
    }
    
    switch absNumber % 10 {
    case 1:
        return "st"
    case 2:  
        return "nd"
    case 3:
        return "rd"
    default:
        return "th"
    }
}

// Ordinalize the number by adding the Ordinal to the number.
func Odinalize(number int) string {
    return strconv.Itoa(number) + Ordinal(number)
}