// Package clock implements a simple library to handle time without dates
package clock

import (
	"fmt"
)

const testVersion = 3

// A Clock represets the clock with hour and minutes
type Clock struct {
	h, m int
}

// Create a new clock representing a clock with hours and minutes
func New(hour, minute int) Clock {
	hour = (hour + minute/60) % 24
	minute = minute % 60
	if minute < 0 {
		hour--
		minute = minute + 60
	}
	if hour < 0 {
		hour = hour + 24
	}
	return Clock{hour, minute}
}

// Return a string representing the actual state of the clock. pattern: hh:mm
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add an quantity of minutes on the clock.
// Negative numbers will back the clock
func (c Clock) Add(minutes int) Clock {
	var hour = (c.h + (c.m+minutes)/60) % 24
	minutes = (minutes + c.m) % 60
	if minutes < 0 {
		hour--
		minutes = minutes + 60
	}
	if hour < 0 {
		hour = hour + 24
	}
	return Clock{hour, minutes}
}
