// STATUS: Successfully passed.

package clock

import "fmt"

// Globally defined consts for general time-keeping
const hours = 24
const minutes = 60
const day = minutes * hours

// Structure to hold Clock instance logic
type Clock struct {
	minutes int
}

// New() Instantiates and adds time to new Clock{} instance
func New(hour, minute int) Clock {
	return Clock{0}.Add(hour*minutes + minute)
}

// Add() adds time incrementally to a Clock{} instance
func (clock Clock) Add(minutes int) Clock {
	clock.minutes += minutes

	if clock.minutes < 0 {
		if clock.minutes < day*-1 {
			clock.minutes = day + (clock.minutes % day)
		} else {
			clock.minutes = day + clock.minutes
		}
	} else if clock.minutes > day {
		clock.minutes %= day
	}
	return clock
}

// String() stringifies the Clock{} instance's time output to be more readable
func (clock Clock) String() string {
	mins := clock.minutes
	hrs := 0
	hrs = mins / minutes
	if hrs == 24 {
		hrs = 0
	}
	mins %= minutes
	return fmt.Sprintf("%02d:%02d", hrs, mins)
}
