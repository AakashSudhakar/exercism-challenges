// STATUS: Successfully passed.

package grains

import (
	"errors"
	"math"
)

// Square() squares a given input integer or returns a custom Error
func Square(integer int) (uint64, error) {
	if integer < 1 || integer > 64 {
		return 0, errors.New("ERROR: Invalid input argument for integer-to-square.")
	}
	return uint64(math.Pow(2, float64(integer-1))), nil
}

// Total() sums all squares in range [1, 64]
func Total() uint64 {
	var sumTotal uint64

	for iterator := 1; iterator <= 64; iterator++ {
		subexp, _ := Square(iterator)
		sumTotal += subexp
	}
	return sumTotal
}
