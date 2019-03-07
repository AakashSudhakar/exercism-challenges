// STATUS: Successfully passed.
package diffsquares

// Difference() calculates the difference in both square expressions
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

// SquareOfSums() squares the sum of all integers in the range [1, n]
func SquareOfSums(n int) int {
	return n * n * ((n * n) + (2 * n) + 1) / 4
}

// SumOfSquares() sums all the squares of every integer in the range [1, n]
func SumOfSquares(n int) int {
	return n * (n + 1) * ((2 * n) + 1) / 6
}
