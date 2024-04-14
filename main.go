package main

import (
	"fmt"
	"github.com/perbu/rgb-weather/strip"
	"github.com/perbu/rgb-weather/yr"
	"os"
	"time"
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
	for {
		f, err := yr.GetForecast(8)
		invalidate := time.Now().Add(1 * time.Hour)
		if err != nil {
			return fmt.Errorf("could not get forecast: %w", err)
		}
		strp := f2f(f)
		strp.Run(invalidate)
	}
}

func f2f(yr yr.Forecast) strip.Forecast {
	f := strip.Forecast{
		Strip: make([]strip.RGBColor, len(yr.Properties.Timeseries)*10),
	}
	f.Hours = make([]*strip.Hour, len(yr.Properties.Timeseries))
	for i, ts := range yr.Properties.Timeseries {
		h := f.NewHour(i, 10)
		h.CloudCover = ts.Data.Instant.Details.CloudAreaFraction
		h.WindSpeed = ts.Data.Instant.Details.WindSpeed
		h.Precipitation = ts.Data.Next_1_hours.Details.PrecipitationAmount
		// fmt.Println("CloudCover:", h.CloudCover, "WindSpeed:", h.WindSpeed, "Precipitation:", h.Precipitation)
		f.Hours[i] = h
	}
	return f
}
