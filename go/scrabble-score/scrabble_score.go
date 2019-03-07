// STATUS: Successfully passed.
package scrabble

import "strings"

// Global variable storing key-value pairs of letters and associated integer scores
var scrabbleMap = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2,
	'B': 3, 'C': 3, 'M': 3, 'P': 3,
	'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5,
	'J': 8, 'X': 8,
	'Q': 10, 'Z': 10,
}

// Global function to calculate total score of word
func Score(word string) int {
	// Initializing score integer
	score := 0

	// Maps over letters in word and accumulates score per letter
	for _, letter := range strings.ToUpper(word) {
		score += scrabbleMap[letter]
	}
	return score
}
