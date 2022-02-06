package clock

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
)

type clockView struct {
	textView *tview.TextView
}

func (c *clockView) update(app *tview.Application) {
	for {
		now := time.Now()
		date := now.Format("Mon Jan 2 2006")
		str := getTimeInSymbols(now)

		fmt.Fprint(c.textView, fmt.Sprintf("%s\n\n%s", str, date))
		app.Draw()

		time.Sleep(1 * time.Second)
		c.textView.Clear()
	}
}

func NewClockView(app *tview.Application) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetTextAlign(tview.AlignCenter)

	view := &clockView{
		textView: textView,
	}

	go view.update(app)

	return view.textView
}

func getTimeInSymbols(t time.Time) string {
	str := t.Format("15:04:05")

	var asciiCharacters [][]string
	for _, char := range str {
		asciiCharacters = append(asciiCharacters, characters[char])
	}

	mergedCharacters := []string{"", "", "", "", ""}
	for char := 0; char < len(asciiCharacters); char++ {
		for row := 0; row < len(asciiCharacters[char]); row++ {
			mergedCharacters[row] = mergedCharacters[row] + asciiCharacters[char][row] + " "
		}
	}

	return strings.Join(mergedCharacters, "\n")
}
