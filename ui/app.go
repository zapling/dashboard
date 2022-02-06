package ui

import (
	"github.com/zapling/dashboard/state"
	"github.com/zapling/dashboard/ui/pages"
)

// Start setups the interface and renders the application
func Start() error {
	state := state.NewUIState()

	landing := pages.NewLandingPage(state)
	state.Pages.AddPage(pages.PAGE_LANDING, landing, true, true)

	calendar := pages.NewCalendarPage(state)
	state.Pages.AddPage(pages.PAGE_CALENDAR, calendar, true, false)

	return state.App.SetRoot(state.Pages, true).Run()
}
