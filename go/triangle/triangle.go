// STATUS: Pending.

package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota
	Equ = iota
	Iso = iota
	Sca = iota
)

// KindFromSides() identifies a triangle type.
func KindFromSides(a, b, c float64) Kind {
	edges := []float64{a, b, c}
	sort.Sort(sort.Float64Slice(edges))

	switch {
	case math.IsNaN(edges[0]) || edges[0] <= 0 || edges[2] == math.Inf(1) || edges[1]+edges[0] <= edges[2]:
		return NaT
	case edges[0] == edges[1] && edges[1] == edges[2]:
		return Equ
	case edges[0] == edges[1] || edges[1] == edges[2]:
		return Iso
	}
	return Sca
}
