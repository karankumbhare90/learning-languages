package clock

import "fmt"

// Clock represents a time without a date
type Clock struct {
	hour   int
	minute int
}

// normalize ensures that hour and minute are within proper ranges
func normalize(h, m int) (int, int) {
	totalMinutes := h*60 + m

	// wrap around 24 hours = 1440 minutes
	totalMinutes = ((totalMinutes % 1440) + 1440) % 1440

	h = totalMinutes / 60
	m = totalMinutes % 60

	return h, m
}

// New creates a new Clock
func New(h, m int) Clock {
	h, m = normalize(h, m)
	return Clock{hour: h, minute: m}
}

// Add adds minutes to the clock
func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

// Subtract subtracts minutes from the clock
func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

// String formats the clock as HH:MM
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
