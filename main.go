package main

import (
	"github.com/zapling/dashboard/ui"
)

func main() {
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
