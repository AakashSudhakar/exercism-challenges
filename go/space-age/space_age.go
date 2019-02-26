// space_age.go (STATUS: In progress.)

package space

// Age function returns converted age from seconds to planetary age.
func Age(seconds float64, planet string) float64 {
	var Sec2Hr, Hr2Day, Day2Year = 3600.00, 24.00, 365.25
	var PlanetaryMap = map[string]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	var PlanetaryAge = float64(PlanetaryMap[planet])
	return float64(PlanetaryAge / (Sec2Hr * Hr2Day * Day2Year))
}
