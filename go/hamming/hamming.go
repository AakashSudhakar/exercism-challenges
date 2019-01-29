package hamming

import "errors"

// Function calculates Hamming distance across two genomes
func Distance(a, b string) (int, error) {
	var HammingError error
	var HammingDistance int
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
