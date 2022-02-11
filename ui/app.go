package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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

	// tmp

	taskPriority := tview.NewTextView()
	taskPriority.SetDynamicColors(true)
	taskPriority.SetBackgroundColor(tcell.NewHexColor(2565927))
	taskPriority.SetBorderPadding(1, 1, 1, 1)

	fmt.Fprint(taskPriority, `[::b]Top priority[-:-:-]
	[aqua::b]t[-:-:-] -> [orange]Today's tasks[-]
	[aqua::b]w[-:-:-] -> [orange]Weekly summary[-]`)

	taskCapture := tview.NewTextView()
	taskCapture.SetDynamicColors(true)
	taskCapture.SetBackgroundColor(tcell.NewHexColor(2565927))
	taskCapture.SetBorderPadding(1, 1, 1, 1)

	fmt.Fprint(taskCapture, `[::b]Capture[-:-:-]
	[aqua::b]c[-:-:-] -> [orange]Capture new task[-]
	`)

	taskView := tview.NewFlex()
	taskView.SetDirection(tview.FlexColumn)
	taskView.AddItem(taskCapture, 0, 1, false)
	taskView.AddItem(taskPriority, 0, 1, false)
	taskView.AddItem(nil, 0, 1, false)

	taskView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyESC {
			uiState.Pages.HidePage("task_view")
		}

		return event
	})

	taskPage := tview.NewFlex().SetDirection(tview.FlexRow)
	taskPage.AddItem(nil, 0, 3, false)
	taskPage.AddItem(taskView, 0, 1, true)

	uiState.Pages.AddPage("task_view", taskPage, true, false)

	return uiState.App.SetRoot(uiState.Pages, true).Run()
}
