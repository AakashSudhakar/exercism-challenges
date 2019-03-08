// STATUS: Pending

package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[string]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA struct {
	sequence string
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (dna DNA) Counts() *Histogram {
	countA, err := dna.Count("A")
	countT, err := dna.Count("T")
	countG, err := dna.Count("G")
	countC, err := dna.Count("C")

	if err != nil {
		fmt.Println(err)
	}
	return &Histogram{"A": countA, "T": countT, "G": countG, "C": countC}
}

// Count() increments count on each individual nucleotide entry
func (dna DNA) Count(nucleotide string) (int, error) {
	nucleotides := map[string]bool{"A": true, "T": true, "G": true, "C": true}

	if !nucleotides[nucleotide] {
		return 0, fmt.Errorf("INVALID NUCLEOTIDE ENTRY: %s", nucleotide)
	}

	cumsum := 0
	for _, value := range dna.sequence {
		if nucleotide == string(value) {
			cumsum++
		}
	}
	return cumsum, nil
}
