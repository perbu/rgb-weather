package main

import (
	"fmt"
	"github.com/perbu/rgb-weather/strip"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	f := strip.GenerateForecast(12, 12)
	f.Run()
	return nil
}
