package main

import (
	"fmt"
	"github.com/perbu/rgb-weather/strip"
	"github.com/perbu/rgb-weather/yr"
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
	// f := strip.GenerateForecast(12, 12)
	f, err := yr.GetForecast(8)
	if err != nil {
		return fmt.Errorf("could not get forecast: %w", err)
	}
	strp := f2f(f)
	strp.Run()
	return nil
}

func f2f(yr yr.Forecast) strip.Forecast {
	f := strip.Forecast{
		Strip: make([]strip.RGBColor, len(yr.Properties.Timeseries)*10),
		Ticks: 0,
	}
	f.Hours = make([]*strip.Hour, len(yr.Properties.Timeseries))
	for i, ts := range yr.Properties.Timeseries {
		h := f.NewHour(i, 10)
		h.CloudCover = ts.Data.Instant.Details.CloudAreaFraction
		fmt.Println("CloudCover:", h.CloudCover)
		h.WindSpeed = ts.Data.Instant.Details.WindSpeed
		h.Precipitation = ts.Data.Next_1_hours.Details.PrecipitationAmount
		f.Hours[i] = h
	}
	return f
}
