// Package clock implements a simple library to handle time without dates
package clock

import (
	"fmt"
)

const testVersion = 3

// Constants to represent the fixed number of hours in a day and minutes in a hour
const (
	dayInHours    = 24
	hourInMinutes = 60
)

// A Clock represets the clock with hour and minutes
type Clock struct {
	h, m int
}

// Create a new clock representing a clock with hours and minutes
func New(hour, minute int) Clock {
	hour = (hour + minute/hourInMinutes) % dayInHours
	minute = minute % hourInMinutes
	if minute < 0 {
		hour--
		minute = minute + hourInMinutes
	}
	if hour < 0 {
		hour = hour + dayInHours
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
	return New(c.h, c.m+minutes)
}
