package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/zapling/dashboard/ui"
)

func main() {
	if err := ui.Start(); err != nil {
		panic(err)
	}
}
