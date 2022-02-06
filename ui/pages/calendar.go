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
	now := time.Now()

	monthView := calendar.NewMonthView(now.Year(), int(now.Month()))

	fmt.Fprint(monthView, "Calendar page, ESC to go back")

	monthView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyESC {
			state.Pages.HidePage(PAGE_CALENDAR)
			state.Pages.SwitchToPage(PAGE_LANDING)
			return event
		}

		return event
	})

	return monthView
}
