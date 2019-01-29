package hamming

import "errors"

// Function calculates Hamming distance across two genomes
func Distance(a, b string) (int, error) {
	var HammingError error
	var HammingDistance int

	// Checks if there is length mismatch across genomes
	// If length matches, checks for letter mismatches
	// If letter mismatch, increment Hamming distance
	if len(a) != len(b) {
		HammingError = errors.New("Length mismatch")
	} else {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				HammingDistance++
			}
		}
	}
	return HammingDistance, HammingError
}
