package clock

import (
	"fmt"
	"math"
)

// Clock type
type Clock struct {
	h, m int
}

// New builds a new Clock type
func New(h, m int) Clock {

	hours, minutes := getHoursMinutes(h, m)

	return Clock{h: hours, m: minutes}
}

// String outputs a string clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add minutes to Clock
func (c Clock) Add(m int) Clock {

	hours, minutes := getHoursMinutes(c.h, (c.m + m))
	c.h = hours
	c.m = minutes
	return c
}

// Subtract minutes to Clock
func (c Clock) Subtract(m int) Clock {
	hours, minutes := getHoursMinutes(c.h, (c.m - m))
	c.h = hours
	c.m = minutes
	return c
}

func getHoursMinutes(h, m int) (int, int) {
	minutes := m % 60

	if minutes < 0 {
		minutes = minutes + 60
	}

	hours := int(m / 60)

	if m < 0 {
		hours = int(math.Ceil(float64(m)/-60.0) * -1)
	}

	hours = (hours + h) % 24
	if hours < 0 {
		hours = hours + 24
	}

	return hours, minutes
}
