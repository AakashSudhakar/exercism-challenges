// STATUS: Pending.

package leap

// IsLeapYear() checks if a year is a leap year or not. Returns a boolean indicating as such.
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	default:
		return false
	}
}
