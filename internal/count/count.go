package count

import (
	"math"
)

// CountToInt returns a slice of integers from 0 to the absolute value of input.
func CountToInt(input float64) []int {
	countTo := int(math.Abs(input))
	result := make([]int, 0, countTo+1)
	for i := 0; i <= countTo; i++ {
		result = append(result, i)
	}
	return result
}
