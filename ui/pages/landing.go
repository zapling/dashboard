package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zapling/dashboard/clock"
	"github.com/zapling/dashboard/state"
)

const PAGE_LANDING = "landing_page"

func NewLandingPage(state *state.UIState) tview.Primitive {
	var clockHeight = 8
	var pageWidth = 60

	emptyTextView := tview.NewTextView()
	clockView := clock.NewClockView(state.App)

	menuTextView := tview.NewTextView()
	menuTextView.SetTextAlign(tview.AlignCenter)
	printMenuText(menuTextView)

	menuTextView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'c' {
			state.Pages.HidePage(PAGE_LANDING)
			state.Pages.SwitchToPage(PAGE_CALENDAR)
			return event
		}

		return event
	})

	pageRows := tview.NewFlex().SetDirection(tview.FlexRow)
	pageRows.AddItem(emptyTextView, 0, 1, false)
	pageRows.AddItem(clockView, clockHeight, 1, false)
	pageRows.AddItem(menuTextView, 0, 1, true)

	page := tview.NewFlex().SetDirection(tview.FlexColumn)
	page.AddItem(emptyTextView, 0, 1, false)
	page.AddItem(pageRows, pageWidth, 1, true)
	page.AddItem(emptyTextView, 0, 1, false)

	return page
}

func printMenuText(tv *tview.TextView) {
	githubNotifications := getNumGithubNotifications()
	gitlabNotifications := getNumGitlabNotifications()

	fmt.Fprint(tv, fmt.Sprintf(" %d  %d\n ", githubNotifications, gitlabNotifications))
	fmt.Fprint(tv, "\n  Calendar        c\n\n")
}

func getNumGithubNotifications() int {
	token := os.Getenv("DASHBOARD_GITHUB_TOKEN")
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/notifications", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth("zapling", token)

	res, _ := http.DefaultClient.Do(req)

	var body []struct{ ID string }
	json.NewDecoder(res.Body).Decode(&body)

	return len(body)
}

func getNumGitlabNotifications() int {
	token := os.Getenv("DASHBOARD_GITLAB_TOKEN")
	req, _ := http.NewRequest(http.MethodGet, "https://gitlab.zimpler.com/api/v4/todos", nil)
	req.Header.Add("PRIVATE-TOKEN", token)

	res, _ := http.DefaultClient.Do(req)

	var body []struct{ ID string }
	json.NewDecoder(res.Body).Decode(&body)

	return len(body)
}
