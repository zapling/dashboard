package calendar

import "time"

type event struct {
	id        string
	text      string
	color     string
	startTime time.Time
	stopTime  time.Time
}

func GetEventsByDates(dates []time.Time) map[time.Time][]event {
	dateEventsMap := make(map[time.Time][]event)

	for _, date := range dates {
		events := GetEventsByDate(date)
		if events == nil {
			continue
		}

		dateEventsMap[date] = events
	}

	return dateEventsMap
}

func GetEventsByDate(d time.Time) []event {
	var events []event

	// TODO: fetch events proper

	// single day event
	if d.Day() == 13 && int(d.Month()) == 2 {
		events = append(events, event{
			text:      " Jacob födelsedag",
			color:     "blue::b",
			startTime: date(2022, 2, 13),
			stopTime:  date(2022, 2, 13),
		})
	}

	// multi-day event (same week)
	if (d.Day() == 8 || d.Day() == 9) && int(d.Month()) == 2 {
		events = append(events, event{
			text:      "  Research day",
			color:     "black:orange",
			startTime: date(2022, 2, 8),
			stopTime:  date(2022, 2, 9),
		})
	}

	// multi-day event (different weeks)
	if (d.Day() == 20 || d.Day() == 21) && int(d.Month()) == 2 {
		events = append(events, event{
			text:      "Multi lines event",
			color:     "black:orange",
			startTime: date(2022, 2, 8),
			stopTime:  date(2022, 2, 9),
		})
	}

	return events
}
