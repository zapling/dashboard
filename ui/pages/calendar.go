package pages

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zapling/dashboard/state"
)

const PAGE_CALENDAR = "calendar_page"

func NewCalendarPage(state *state.UIState) tview.Primitive {
	calendarTextView := tview.NewTextView()

	fmt.Fprint(calendarTextView, "Calendar page, ESC to go back")

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
