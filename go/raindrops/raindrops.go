// raindrops.go (STATUS: Pending)

package raindrops

import "strconv"

// Convert function returns a select string response based on numerical patterns of its input param.
// NOTE: Simple solution -- consider using Map object with helper function for optimal quality?
func Convert(number int) string {
	var PlingFactor, PlangFactor, PlongFactor int = 3, 5, 7
	var PlingEl, PlangEl, PlongEl string = "Pling", "Plang", "Plong"
	var Raindrop string

	if number%PlingFactor == 0 {
		// Pling
		Raindrop += PlingEl
	}
	if number%PlangFactor == 0 {
		// Plang
		Raindrop += PlangEl
	}
	if number%PlongFactor == 0 {
		// Plong
		Raindrop += PlongEl
	}
	if Raindrop == "" {
		Raindrop = strconv.Itoa(number)
	}
	return Raindrop
}
