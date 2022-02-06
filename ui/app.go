package ui

import (
	"github.com/zapling/dashboard/state"
	"github.com/zapling/dashboard/ui/pages"
)

// Start setups the interface and renders the application
func Start() error {
	uiState := state.NewUIState()

	state.GlobalUIState = uiState

	landing := pages.NewLandingPage(uiState)
	uiState.Pages.AddPage(pages.PAGE_LANDING, landing, true, true)

	calendar := pages.NewCalendarPage(uiState)
	uiState.Pages.AddPage(pages.PAGE_CALENDAR, calendar, true, false)

	return uiState.App.SetRoot(uiState.Pages, true).Run()
}
