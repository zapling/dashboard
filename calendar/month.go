package calendar

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
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
// └──────────┴───────────┴───────────┴───────────┴───────────┴───────────┴───────────┘

type monthView struct {
	textView *tview.TextView
	month    time.Time
	days     []time.Time
}

func (m *monthView) NextMonth() {
	m.changeMonth(m.month.AddDate(0, 1, 0))
}

func (m *monthView) PrevMonth() {
	m.changeMonth(m.month.AddDate(0, -1, 0))
}

func (m *monthView) CurrentMonth() {
	m.changeMonth(time.Now())
}

func (m *monthView) changeMonth(t time.Time) {
	m.month = t
	m.days = getMonthDays(t)
	m.textView.Clear()
	m.render()
}

func (m *monthView) getInputHandler(exitCallback func()) func(event *tcell.EventKey) *tcell.EventKey {
	return func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '>' {
			m.NextMonth()
			return event
		}

		if event.Rune() == '<' {
			m.PrevMonth()
			return event
		}

		if event.Rune() == '.' {
			m.CurrentMonth()
			return event

		}

		if event.Key() == tcell.KeyESC {
			exitCallback()
			return event
		}
		return event
	}
}

func (m *monthView) render() {
	renderMonthCalendar(m)
}

func NewMonthView(month time.Time, exitCallback func()) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)

	view := &monthView{
		textView: textView,
		month:    month,
		days:     getMonthDays(month),
	}

	textView.SetInputCapture(view.getInputHandler(exitCallback))

	renderMonthCalendar(view)

	return view.textView
}

func renderMonthCalendar(m *monthView) {
	header := fmt.Sprintf(`
                                  %s %d / %d
┌───────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
│  Monday   │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ [blue]Saturday[-]  │  [red]Sunday[-]   │
├───────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
`, m.month.Month().String(), m.month.Year(), m.month.Month())

	// render header
	fmt.Fprint(m.textView, header)

	numWeeks := len(m.days) / 7
	for i := 0; i < numWeeks; i++ {
		weekDays := m.days[i*7 : 7+(i*7)]

		var dateRow string
		for _, day := range weekDays {
			color := ":"

			if day.Month() != m.month.Month() {
				color = "grey"
			} else if day.Weekday() == time.Saturday {
				color = "blue"
			} else if day.Weekday() == time.Sunday {
				color = "red"
			}

			if dateEqual(day, time.Now()) {
				color = "darkgreen:green"
			}

			dateRow = dateRow + appendSpaces(fmt.Sprintf("│ [%s]%2d[-:-:-]", color, day.Day()), 8)
		}

		fmt.Fprintf(m.textView, dateRow+"│\n")

		// render 3 empty rows for now
		emptyRow := appendSpaces("│", 11)
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

func getMonthDays(month time.Time) []time.Time {
	var days []time.Time

	firstDay := date(month.Year(), int(month.Month()), 1)
	lastDay := date(month.Year(), int(month.Month())+1, 0)

	// calculate how many days we need to go back to get a full week
	startOffset := 0
	if firstDay.Weekday() != time.Monday {
		subtract := int(firstDay.Weekday())
		if firstDay.Weekday() == time.Sunday {
			subtract = 7
		}

		startOffset = int(time.Monday) - subtract
	}

	// calculate how may days forward we need to go until we get a full week
	stopDay := lastDay
	if stopDay.Weekday() != time.Sunday {
		daysUntilCompleteWeek := int(time.Saturday) + 1 - int(lastDay.Weekday())
		stopDay = lastDay.AddDate(0, 0, daysUntilCompleteWeek)
	}

	for {
		day := date(firstDay.Year(), int(firstDay.Month()), 1+startOffset)

		if day.After(stopDay) {
			break
		}

		days = append(days, day)

		startOffset++
	}

	return days
}
