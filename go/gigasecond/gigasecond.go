// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	giga := time.Duration((math.Pow(10.0, 9.0)))
	return t.Add(giga * time.Second)
}
