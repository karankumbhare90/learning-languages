package meetup

import "time"

// WeekSchedule defines the type of schedule (first, second, etc.)
type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the day of the month for the given schedule.
func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// Start from the first day of the month
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	switch wSched {
	case Teenth:
		// Teenth days: 13th–19th
		for d := 13; d <= 19; d++ {
			t = time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == wDay {
				return d
			}
		}

	case Last:
		// Get last day of month by going to next month’s 0th day
		lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC)
		for d := lastDay.Day(); d > 0; d-- {
			t = time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == wDay {
				return d
			}
		}

	default:
		// Map week schedule to offset (0 for First, 1 for Second, etc.)
		offset := int(wSched)
		// Find first occurrence of that weekday in month
		for d := 1; d <= 7; d++ {
			t = time.Date(year, month, d, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == wDay {
				return d + (offset * 7)
			}
		}
	}

	panic("Invalid schedule or weekday combination")
}
