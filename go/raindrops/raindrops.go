// raindrops.go (STATUS: Passed.)

package raindrops

import "strconv"

// Convert function returns a select string response based on numerical patterns of its input param.
// NOTE: Uses Map object with helper function for optimal code quality.
func Convert(number int) string {
	var Raindrop string
	var Plelements = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	// Iterates through key-value indices in Pl-Elements table
	for key, value := range Plelements {
		if number%key == 0 {
			Raindrop += value
		}
	}

	// Final check to convert blank string object appropriately
	if Raindrop == "" {
		return strconv.Itoa(number)
	}
	return Raindrop
}
