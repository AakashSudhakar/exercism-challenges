// raindrops.go (STATUS: Pending)

package raindrops

import "strconv"

// Convert function returns a select string response based on numerical patterns of its input param.
// NOTE: Simple solution -- consider using Map object with helper function for optimal quality?
func Convert(number int) string {
	var Raindrop string
	var Plelements = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	for key, value := range Plelements {
		if number%key == 0 {
			Raindrop += value
		}
	}

	if Raindrop == "" {
		return strconv.Itoa(number)
	}
	return Raindrop
}
