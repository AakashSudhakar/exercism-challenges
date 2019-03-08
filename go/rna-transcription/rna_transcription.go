// STATUS: Successfully submitted.

package strand

import "strings"

// ToRNA() converts a DNA input string into an RNA output transcribed string.
func ToRNA(strDNA string) string {
	return transcriptionMap.Replace(strDNA)
}

var transcriptionMap = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"G", "C",
	"C", "G",
)
