package pages

import (
	"time"

	"github.com/rivo/tview"
	"github.com/zapling/dashboard/calendar"
	"github.com/zapling/dashboard/state"
)

const PAGE_CALENDAR = "calendar_page"

func NewCalendarPage(state *state.UIState) tview.Primitive {
	now := time.Now()

	monthView := calendar.NewMonthView(now, func() {
		state.Pages.HidePage(PAGE_CALENDAR)
		state.Pages.SwitchToPage(PAGE_LANDING)
	})

	return monthView
}
