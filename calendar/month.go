package calendar

import (
	"time"
)

// GetMonth returns dates for 5 full weeks
func GetMonth(t time.Time) []time.Time {
	var month []time.Time

	firstDay := date(t.Year(), int(t.Month()), 1)

	offset := int(time.Monday) - int(firstDay.Weekday())

	week := 0
	for {
		if week == 5 {
			break
		}

		day := date(firstDay.Year(), int(firstDay.Month()), 1+offset)

		month = append(month, day)

		// if we added a sunday we are done with the current week
		if day.Weekday() == time.Sunday {
			week++
		}

		offset++
	}

	return month
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
