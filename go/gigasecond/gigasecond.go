// Package gigasecond implements a simple library to calculate the 1 Gs anniversary.
// A gigasecond is one billion (10**9) seconds.
package gigasecond

import "time"

// Test version.
const testVersion = 4

// AddGigasecond returns the time.Time + 1e9 seconds.
func AddGigasecond(date time.Time) time.Time {
	var d = time.Duration(1e9 * time.Second)
	return date.Add(d)
}
