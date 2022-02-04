package main

import (
	"fmt"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	BACKGROUND_COLOR = 2565927
)

var app = tview.NewApplication()

func main() {
	box := tview.NewBox().SetBorder(false).SetTitle("Hello, world!")
	box.SetBackgroundColor(tcell.NewHexColor(BACKGROUND_COLOR))

	textBox := tview.NewTextView()
	textBox.SetBackgroundColor(tcell.NewHexColor(BACKGROUND_COLOR))

	height := 9
	width := 60

	layout := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(box, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(textBox, height, 1, true).
			AddItem(nil, 0, 1, false),
			width, 1, false).
		AddItem(nil, 0, 1, true)

	go printClock(app, textBox)

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func printClock(app *tview.Application, primitive *tview.TextView) {
	for true {
		now := time.Now()
		displayTime := now.Format("15:04:05")
		date := now.Format("Mon Jan 2 2006")

		// ascii time
		myFigure := figure.NewFigure(displayTime, "big", true)
		str := myFigure.String() + "\n"

		spacesToAdd := 30 - len(date)/2

		for i := 0; i < spacesToAdd; i++ {
			str = str + " "
		}

		// append date
		str = str + date

		fmt.Fprint(primitive, str)
		app.Draw()

		time.Sleep(1 * time.Second)
		primitive.Clear()
	}
}
