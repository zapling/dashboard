package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/zapling/dashboard/asciiart"
	"github.com/zapling/dashboard/state"
)

func NewLandingPage(state *state.UIState) tview.Primitive {
	var clockHeight = 8
	var pageWidth = 60

	emptyTextView := tview.NewTextView()

	clockTextView := tview.NewTextView()
	clockTextView.SetTextAlign(tview.AlignCenter)
	go updateClock(clockTextView, state.App)

	menuTextView := tview.NewTextView()
	menuTextView.SetTextAlign(tview.AlignCenter)
	printMenuText(menuTextView)

	pageRows := tview.NewFlex().SetDirection(tview.FlexRow)
	pageRows.AddItem(emptyTextView, 0, 1, false)
	pageRows.AddItem(clockTextView, clockHeight, 1, true)
	pageRows.AddItem(menuTextView, 0, 1, false)

	page := tview.NewFlex().SetDirection(tview.FlexColumn)
	page.AddItem(emptyTextView, 0, 1, false)
	page.AddItem(pageRows, pageWidth, 1, false)
	page.AddItem(emptyTextView, 0, 1, false)

	return page
}

func printMenuText(tv *tview.TextView) {
	githubNotifications := getNumGithubNotifications()
	gitlabNotifications := getNumGitlabNotifications()

	fmt.Fprint(tv, fmt.Sprintf(" %d  %d\n ", githubNotifications, gitlabNotifications))
}

func updateClock(tv *tview.TextView, app *tview.Application) {
	for true {
		now := time.Now()
		date := now.Format("Mon Jan 2 2006")

		str := asciiart.GetTime(now) + "\n\n" + date

		fmt.Fprint(tv, str)
		app.Draw()

		time.Sleep(1 * time.Second)
		tv.Clear()
	}
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
