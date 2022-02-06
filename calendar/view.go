package calendar

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

//                                      YYYY / MM
//
// ┌──────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
// │  Monday  │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ Saturday  │  Sunday   │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 01       │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// ├──────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
// │ 1        │ 2         │ 2         │ 2         │ 2         │ 2         │ 2         │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// │          │           │           │           │           │           │           │
// └──────────┴───────────┴───────────┴───────────┴───────────┴───────────┴───────────┘`

type monthView struct {
	textView  *tview.TextView
	viewYear  int
	viewMonth int
	days      []time.Time
}

func (m *monthView) render() {
	header := fmt.Sprintf(`                                      %d / %d
┌───────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
│  Monday   │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ [blue]Saturday[-]  │  [red]Sunday[-]   │
├───────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
`, m.viewYear, m.viewMonth)

	// render header
	fmt.Fprint(m.textView, header)

	numWeeks := len(m.days) / 7
	for i := 0; i < numWeeks; i++ {
		weekDays := m.days[i*7 : 7+(i*7)]

		var dateRow string
		for _, day := range weekDays {
			color := ":"

			if day.Month() != time.Month(m.viewMonth) {
				color = "grey"
			} else if day.Weekday() == time.Saturday {
				color = "blue"
			} else if day.Weekday() == time.Sunday {
				color = "red"
			}

			if dateEqual(day, time.Now()) {
				color = "darkgreen:green"
			}

			dateRow = dateRow + m.appendSpaces(fmt.Sprintf("│ [%s]%2d[-:-:-]", color, day.Day()), 8)
		}

		fmt.Fprintf(m.textView, dateRow+"│\n")

		// render 3 empty rows for now
		emptyRow := m.appendSpaces("│", 11)
		for i := 0; i < 3; i++ {
			for y := 0; y < 7; y++ {
				fmt.Fprint(m.textView, emptyRow)
			}
			fmt.Fprint(m.textView, "│\n")
		}

		// render closing line

		leftCorner := "├"
		rightCorner := "┤"
		middle := "┼"

		if i == numWeeks-1 {
			leftCorner = "└"
			rightCorner = "┘"
			middle = "┴"
		}

		closingLine := fmt.Sprintf(
			`%[1]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[3]s`,
			leftCorner, middle, rightCorner,
		)

		fmt.Fprint(m.textView, closingLine+"\n")
	}

}

func (m *monthView) appendSpaces(str string, amount int) string {
	for i := 0; i < amount; i++ {
		str = str + " "
	}

	return str
}

func NewMonthView(year, month int) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)

	view := &monthView{
		textView:  textView,
		viewYear:  year,
		viewMonth: month,
		days:      GetMonthView(year, month),
	}

	view.render()

	return view.textView
}
