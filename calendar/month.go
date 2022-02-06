package calendar

import (
	"fmt"
	"time"
)

// GetMonth returns dates for 5 full weaks
func GetMonth(t time.Time) string {
	var str string

	firstDay := date(t.Year(), int(t.Month()), 1)

	offset := int(time.Monday) - int(firstDay.Weekday())

	week := 0
	for {
		if week == 5 {
			break
		}

		day := date(firstDay.Year(), int(firstDay.Month()), 1+offset)

		str = fmt.Sprintf("%s %d ", str, day.Day())

		// if we added a sunday we are done with the current week
		if day.Weekday() == time.Sunday {
			str = str + "\n"
			week++
		}

		offset++
	}

	return str
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
