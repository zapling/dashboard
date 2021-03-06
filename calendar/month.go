package calendar

import (
	"fmt"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const calendarHeader string = `
%s %d / %d
┌───────────┬───────────┬───────────┬───────────┬───────────┬───────────┬───────────┐
│  Monday   │  Tuesday  │ Wednsday  │ Thursday  │  Friday   │ [blue]Saturday[-]  │  [red]Sunday[-]   │
├───────────┼───────────┼───────────┼───────────┼───────────┼───────────┼───────────┤
`

const weekDivider string = `%[1]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[2]s───────────%[3]s`

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

func NewMonthView(month time.Time, exitCallback func()) *tview.TextView {
	textView := tview.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetTextAlign(tview.AlignCenter)

	view := &monthView{
		textView:  textView,
		month:     month,
		days:      getMonthDays(month),
		cursorPos: cursorPosition{},
	}

	cursorStartingPos := view.getCursorPosition(time.Now())

	view.cursorPos = cursorStartingPos

	textView.SetInputCapture(view.getInputHandler(exitCallback))
	view.render()

	return view.textView
}

type monthView struct {
	textView  *tview.TextView
	month     time.Time
	days      []time.Time // TODO: rename to daysInView
	cursorPos cursorPosition
}

type cursorPosition struct {
	x int
	y int
}

func (m *monthView) NextMonth() {
	m.changeMonth(m.month.AddDate(0, 1, 0))
}

func (m *monthView) PrevMonth() {
	m.changeMonth(m.month.AddDate(0, -1, 0))
}

func (m *monthView) CurrentMonth() {
	now := time.Now()
	m.changeMonth(now)
	// move cursor after month change, otherwise we won't find the correct day
	m.cursorPos = m.getCursorPosition(now)
	m.render()
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

		// cursor movment
		if event.Rune() == 'j' {
			m.moveCursorPos(m.cursorPos.x, m.cursorPos.y+1)
			return event
		}

		if event.Rune() == 'k' {
			m.moveCursorPos(m.cursorPos.x, m.cursorPos.y-1)
			return event
		}

		if event.Rune() == 'h' {
			m.moveCursorPos(m.cursorPos.x-1, m.cursorPos.y)
			return event
		}

		if event.Rune() == 'l' {
			m.moveCursorPos(m.cursorPos.x+1, m.cursorPos.y)
			return event
		}

		if event.Key() == tcell.KeyESC {
			exitCallback()
			return event
		}
		return event
	}
}

func (m *monthView) getCursorPosition(date time.Time) cursorPosition {
	numWeeks := len(m.days) / 7

	for y := 0; y < numWeeks; y++ {
		weekDays := m.days[y*7 : 7+(y*7)]

		for x := 0; x < len(weekDays); x++ {
			if dateEqual(date, weekDays[x]) {
				return cursorPosition{
					x: x,
					y: y,
				}
			}
		}
	}

	// default on Wednsday week 3 in the view
	// that should be somewhere in the middle, I think :P
	return cursorPosition{x: 3, y: 2}
}

func (m *monthView) moveCursorPos(x, y int) {

	// check if new cursor pos is valid
	maxY := len(m.days)/7 - 1
	if y > maxY {
		m.NextMonth()
		y = 0
	}

	if y < 0 {
		m.PrevMonth()
		maxY = len(m.days)/7 - 1
		y = maxY
	}

	if x < 0 {
		m.PrevMonth()
		x = 6
	}

	if x > 6 {
		m.NextMonth()
		x = 0
	}

	m.cursorPos.x = x
	m.cursorPos.y = y

	m.render()
}

func (m *monthView) render() {
	m.textView.Clear()

	numWeeks := len(m.days) / 7

	calendarRows := m.getCalendarRows()

	// header
	fmt.Fprint(
		m.textView,
		fmt.Sprintf(calendarHeader, m.month.Month().String(), m.month.Year(), m.month.Month()),
	)

	// render weeks divider
	for weekIndex, weekStr := range calendarRows {
		leftCorner := "├"
		rightCorner := "┤"
		middle := "┼"

		lastWeek := false
		if weekIndex == numWeeks-1 {
			leftCorner = "└"
			rightCorner = "┘"
			middle = "┴"
			lastWeek = true
		}

		divider := fmt.Sprintf(weekDivider, leftCorner, middle, rightCorner)

		fmt.Fprint(m.textView, weekStr+"\n")
		fmt.Fprint(m.textView, divider)

		if !lastWeek {
			fmt.Fprint(m.textView, "\n")
		}
	}
}

func (m *monthView) getCalendarRows() []string {
	numWeeks := len(m.days) / 7

	var calendarRows []string
	for y := 0; y < numWeeks; y++ {
		days := m.days[y*7 : 7+(y*7)]

		var weekStr []string
		for x := 0; x < len(days); x++ {
			day := days[x]

			dayStr := m.getCalendarDay(day, x, y)

			weekStr = append(weekStr, strings.Join(dayStr, "\n"))
		}

		tmp := []string{"", "", "", ""}
		for dayId, dayStr := range weekStr {
			for rowId, rowStr := range strings.Split(dayStr, "\n") {
				tmp[rowId] = tmp[rowId] + rowStr
			}

			// append border as the last character on each row if we are on a sunday
			if dayId == 6 {
				for idx := range tmp {
					tmp[idx] = tmp[idx] + "│"
				}
			}

		}

		calendarRows = append(calendarRows, strings.Join(tmp, "\n"))
	}

	return calendarRows
}

func (m *monthView) getCalendarDay(day time.Time, x, y int) []string {
	cursor := ":"

	dayNumberColor := m.getDayNumberColor(day)

	if m.cursorPos.y == y && m.cursorPos.x == x {
		cursor = "-:grey"
		if !dateEqual(day, time.Now()) {
			dayNumberColor = "-:grey"
			if (day.Weekday() == time.Saturday || day.Weekday() == time.Sunday) && day.Month() == m.month.Month() {
				color := m.getDayNumberColor(day)
				dayNumberColor = strings.ReplaceAll(dayNumberColor, "-", color)
			}
		}
	}

	dayStr := fmt.Sprintf(
		"│[%s] [%s]%2d[-:-][%s]        [-:-]\n"+
			"│[%s]           [-:-]\n"+
			"│[%s]           [-:-]\n"+
			"│[%s]           [-:-]",
		cursor, dayNumberColor, day.Day(), cursor, cursor, cursor, cursor,
	)

	return strings.Split(dayStr, "\n")
}

func (m *monthView) getDayNumberColor(day time.Time) string {
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

	return color
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
