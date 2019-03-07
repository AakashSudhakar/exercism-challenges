// STATUS: Submitted successfully.
package luhn

import (
	"strconv"
	"strings"
)

// AddCheck() uses Luhn algorithm to calculate numerical check value
func AddCheck(inputNumStr string) string {
	cumsum := Sum(inputNumStr + "0")
	numCheck := cumsum * 9 % 10
	return inputNumStr + strconv.Itoa(numCheck)
}

// Valid() verifies validity of checked number
func Valid(inputNumStr string) bool {
	cumsum := Sum(inputNumStr)
	return cumsum > 0 && cumsum%10 == 0
}

// Sum() sums the input number (string type)
func Sum(inputNumStr string) (cumsum int) {
	digits := strings.Replace(inputNumStr, " ", "", -1)

	for iterator := 0; iterator < len(digits); iterator++ {
		digit, _ := strconv.Atoi(digits[iterator : iterator+1])
		if (len(digits)-iterator)%2 == 0 {
			digit_doubled := digit * 2
			cumsum += (digit_doubled / 10) + (digit_doubled % 10)
		} else {
			cumsum += digit
		}
	}
	return
}
