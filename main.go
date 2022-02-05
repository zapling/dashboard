package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zapling/dashboard/asciiart"
)

const (
	BACKGROUND_COLOR = 2565927
)

var app = tview.NewApplication()

func main() {
	clockBox := tview.NewTextView()
	clockBox.SetBackgroundColor(tcell.NewHexColor(BACKGROUND_COLOR))
	clockBox.SetTextAlign(tview.AlignCenter)

	statusBox := tview.NewTextView()
	statusBox.SetTextAlign(tview.AlignCenter)
	// statusBox.SetBackgroundColor(tcell.NewHexColor(BACKGROUND_COLOR))

	weather := getWeather()
	if len(weather) > 0 {
		fmt.Fprint(statusBox, getWeather()+"  ")
	}

	fmt.Fprint(statusBox, fmt.Sprintf(" %d ", numGithubNotifications()))
	fmt.Fprint(statusBox, fmt.Sprintf(" %d", numGitlabNotifications()))

	height := 7
	width := 60

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(clockBox, height, 1, true).
			AddItem(statusBox, 0, 2, false),
			width, 1, false).
		AddItem(nil, 0, 1, true)

	go printClock(app, clockBox)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func getWeather() string {
	output, err := exec.Command("weatherapplet").Output()
	if err != nil {
		return ""
	}

	return strings.ReplaceAll(string(output), "\n", "")
}

func printClock(app *tview.Application, primitive *tview.TextView) {
	for true {
		now := time.Now()
		date := now.Format("Mon Jan 2 2006")

		str := asciiart.GetTime(now) + "\n\n" + date

		fmt.Fprint(primitive, str)
		app.Draw()

		time.Sleep(1 * time.Second)
		primitive.Clear()
	}
}

func numGithubNotifications() int {
	token := os.Getenv("DASHBOARD_GITHUB_TOKEN")
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/notifications", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.SetBasicAuth("zapling", token)

	res, _ := http.DefaultClient.Do(req)

	var body []struct{ ID string }
	json.NewDecoder(res.Body).Decode(&body)

	return len(body)
}

func numGitlabNotifications() int {
	token := os.Getenv("DASHBOARD_GITLAB_TOKEN")
	req, _ := http.NewRequest(http.MethodGet, "https://gitlab.zimpler.com/api/v4/todos", nil)
	req.Header.Add("PRIVATE-TOKEN", token)

	res, _ := http.DefaultClient.Do(req)

	var body []struct{ ID string }
	json.NewDecoder(res.Body).Decode(&body)

	return len(body)
}
