package calendar

import (
	"time"
)

func GetMonthView(year, month int) []time.Time {
	var days []time.Time

	firstDay := date(year, month, 1)
	lastDay := date(year, month+1, 0)

	// calculate how many days we need to go back to get a full week
	startOffset := 0
	if firstDay.Weekday() != time.Monday {
		subtract := int(firstDay.Weekday())
		if firstDay.Weekday() == time.Sunday {
			subtract = 7
		}

		startOffset = int(time.Monday) - subtract
	}

	// calculate how may days forward we need to go until we get a full week
	stopDay := lastDay
	if stopDay.Weekday() != time.Sunday {
		daysUntilCompleteWeek := int(time.Saturday) + 1 - int(lastDay.Weekday())
		stopDay = lastDay.AddDate(0, 0, daysUntilCompleteWeek)
	}

	for {
		day := date(firstDay.Year(), int(firstDay.Month()), 1+startOffset)

		if day.After(stopDay) {
			break
		}

		days = append(days, day)

		startOffset++
	}

	return days
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
