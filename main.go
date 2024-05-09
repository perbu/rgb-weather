package main

import (
	"fmt"
	"github.com/perbu/bobblehat/sense/screen"
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
	for {
		f, err := yr.GetForecast(8)
		fmt.Println("Got forecast")
		invalidate := time.Now().Add(1 * time.Hour)
		if err != nil {
			return fmt.Errorf("could not get forecast: %w", err)
		}
		strp, err := f2f(f)
		if err != nil {
			return fmt.Errorf("could not convert forecast: %w", err)
		}
		strp.Run(invalidate)
	}
}

func f2f(yr yr.Forecast) (strip.Forecast, error) {
	dev, err := screen.New()
	if err != nil {
		return strip.Forecast{}, fmt.Errorf("initialize screen device: %w", err)
	}
	f := strip.Forecast{
		Strip: make([]strip.RGBColor, len(yr.Properties.TimeSeries)*10),
		FB:    screen.NewFrameBuffer(),
		Dev:   dev,
	}
	f.Hours = make([]*strip.Hour, len(yr.Properties.TimeSeries))
	for i, ts := range yr.Properties.TimeSeries {
		h := f.NewHour(i, 8)
		h.CloudCover = ts.Data.Instant.Details.CloudAreaFraction
		h.WindSpeed = ts.Data.Instant.Details.WindSpeed
		h.Precipitation = ts.Data.Next1Hours.Details.PrecipitationAmount
		f.Hours[i] = h
		fmt.Printf("Time: %v\n", ts.Time)
		fmt.Printf("CloudCover: %v\n", h.CloudCover)
		fmt.Printf("WindSpeed: %v\n", h.WindSpeed)
		fmt.Printf("Precipitation: %v\n", h.Precipitation)
	}

	return f, nil
}
