package strip

import (
	"fmt"
	"time"
)

const (
	hoursInForecast = 12
	frameRate       = 10
	ledsPerHour     = 10
)

type Forecast struct {
	Hours []*Hour
	Strip []RGBColor
	Ticks int
}

func New(leds int) *Forecast {
	f := &Forecast{
		Strip: make([]RGBColor, leds),
		Ticks: 0,
	}
	f.Hours = make([]*Hour, hoursInForecast)
	for i := range f.Hours {
		f.Hours[i] = f.NewHour(leds, ledsPerHour)
	}
	return f
}

func (f *Forecast) Update() {
	f.Ticks++
	for _, hour := range f.Hours {
		hour.Update(f.Ticks)
	}
}

func (f *Forecast) Display() {
	// Move cursor to 1,1:
	fmt.Print("\x1b[H")
	fmt.Print("\r") // Move cursor to the beginning of the line
	for _, led := range f.Strip {
		fmt.Printf("\x1b[48;2;%d;%d;%dm \x1b[0m", led.R, led.G, led.B)
	}
	fmt.Print("\x1b[0K") // Clear from cursor to the end of the line
	fmt.Println()
}

func (f *Forecast) Run(until time.Time) {
	// create a ticker that ticks every 1/frameRate seconds
	ticker := time.NewTicker(time.Second / frameRate)
	defer ticker.Stop()
	timeout := time.NewTimer(until.Sub(time.Now()))
	for {
		select {
		case <-timeout.C:
			return
		case <-ticker.C:
			f.Update()
			f.Ticks++
			f.Display()
		}

	}

}
