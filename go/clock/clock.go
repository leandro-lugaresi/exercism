// Package clock implements a simple library to handle time without dates.
package clock

import (
	"fmt"
)

const testVersion = 3

// Constants to represent the fixed number of hours in a day and minutes in a hour.
const (
	dayInHours      = 24
	hourInMinutes   = 60
	minutesInOneDay = 1440
)

// A Clock represets the clock using the minutes from 0 to 1440 minutos.
type Clock struct {
	time int
}

// Create a new clock representing a clock with hours and minutes.
func New(hours, minutes int) Clock {
	minutes = (minutes + hours*hourInMinutes) % minutesInOneDay
	if minutes < 0 {
		minutes = minutes + minutesInOneDay
	}
	return Clock{minutes}
}

// Return a string representing the actual state of the clock. pattern: hh:mm.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.time/hourInMinutes, c.time%hourInMinutes)
}

// Add an quantity of minutes on the clock.
// Negative numbers will back the clock.
func (c Clock) Add(minutes int) Clock {
	return New(0, c.time+minutes)
}
