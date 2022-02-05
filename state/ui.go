package state

import "github.com/rivo/tview"

var GlobalUIState *UIState

type UIState struct {
	App *tview.Application
}

// NewUIState creates a new UI state
func NewUIState() *UIState {
	return &UIState{
		App: tview.NewApplication(),
	}
}
