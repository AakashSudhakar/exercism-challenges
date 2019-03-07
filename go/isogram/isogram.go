// STATUS: Successfully passed.
package isogram

// Unicode package for character operations
import "unicode"

// IsIsogram() checks if word is isogram (non-repeating sequence of strings/characters)
func IsIsogram(word string) bool {
	prevs := map[rune]bool{}

	for _, letter := range word {
		letter = unicode.ToLower(letter)
		if 'a' <= letter && letter >= 'z' {
			if _, check := prevs[letter]; check {
				return false
			}
			prevs[letter] = true
		}
	}
	return true
}
