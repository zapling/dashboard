package pages

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zapling/dashboard/calendar"
	"github.com/zapling/dashboard/state"
)

const PAGE_CALENDAR = "calendar_page"

func NewCalendarPage(state *state.UIState) tview.Primitive {
	calendarTextView := tview.NewTextView()

	fmt.Fprint(calendarTextView, getCalendarAscii())
	fmt.Fprint(calendarTextView, "\nCalendar page, ESC to go back")

	calendarTextView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			state.Pages.HidePage(PAGE_CALENDAR)
			state.Pages.SwitchToPage(PAGE_LANDING)
			return event
		}

		return event
	})

	return calendarTextView
}

func getCalendarAscii() string {
	dates := calendar.GetMonth(time.Now())

	var str = `┌───────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
│  Monday   │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ Saturday  │  Sunday   │
├───────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
`

	week := 1
	for _, day := range dates {
		str = str + fmt.Sprintf("│ %02d        ", day.Day())

		if day.Weekday() == time.Sunday {
			str = str + "│"

			left := "├"
			right := "┤"
			middle := "┼"

			if week == 5 {
				left = "└"
				right = "┘"
				middle = "┴"
			}

			str = str + fmt.Sprintf(`
│           │           │           │           │           │           │           │
│           │           │           │           │           │           │           │
│           │           │           │           │           │           │           │
%[1]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[3]s
`, left, middle, right)

			week++
		}
	}

	return str
}

// var calendar = `
//                                      YYYY / MM

// ┌──────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
// │  Monday  │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ Saturday  │  Sunday   │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 01       │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// └──────────┴───────────┴───────────┴───────────┴───────────┴───────────┴───────────┘`
