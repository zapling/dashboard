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

	var calendarHeight = 36
	var pageWidth = 87

	emptyTextView := tview.NewTextView()

	monthView := calendar.NewMonthView(now, func() {
		state.Pages.HidePage(PAGE_CALENDAR)
		state.Pages.SwitchToPage(PAGE_LANDING)
	})

	pageRows := tview.NewFlex().SetDirection(tview.FlexRow)
	pageRows.AddItem(emptyTextView, 0, 1, false)
	pageRows.AddItem(monthView, calendarHeight, 1, true)
	pageRows.AddItem(emptyTextView, 0, 1, false)

	page := tview.NewFlex().SetDirection(tview.FlexColumn)
	page.AddItem(emptyTextView, 0, 1, false)
	page.AddItem(pageRows, pageWidth, 1, true)
	page.AddItem(emptyTextView, 0, 1, false)

	return page
}
