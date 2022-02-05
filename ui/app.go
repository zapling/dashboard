package ui

import (
	"github.com/zapling/dashboard/state"
	"github.com/zapling/dashboard/ui/pages"
)

// Start setups the interface and renders the application
func Start() error {
	state := state.NewUIState()

	landing := pages.NewLandingPage(state)

	return state.App.SetRoot(landing, true).Run()
}
