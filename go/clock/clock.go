package clock

import (
	"fmt"
)

const testVersion = 3

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	hour = (hour + minute/60) % 24
	minute = minute % 60
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (Clock) Add(minutes int) Clock {
	return Clock{1, minutes}
}
