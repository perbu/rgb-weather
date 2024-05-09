package strip

import (
	"fmt"
	"github.com/perbu/bobblehat/sense/screen"
	"time"
)

const (
	hoursInForecast = 8
	frameRate       = 30
	ledsPerHour     = 8
)

type Forecast struct {
	Hours []*Hour
	Strip []RGBColor
	Ticks int
	FB    *screen.FrameBuffer
	Dev   screen.Device
}

// not in use atm, see f2f in main.
func New(leds int) *Forecast {
	f := &Forecast{
		Strip: make([]RGBColor, leds),
		Ticks: 0,
		FB:    screen.NewFrameBuffer(),
	}
	fmt.Printf("FrameBuffer created. Bounds: %v\n", f.FB.Bounds())
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
	for hour := 0; hour < 8; hour++ {
		for led := 0; led < 8; led++ {
			f.FB.SetPixel(led, hour, f.Strip[hour*8+led].PiColor())
		}
	}
	err := f.Dev.Draw(f.FB)
	if err != nil {
		fmt.Printf("Error drawing to device: %v\n", err)
	}

}

func (f *Forecast) Run(until time.Time) {
	// create a ticker that ticks every 1/frameRate seconds
	ticker := time.NewTicker(time.Second / frameRate)
	fmt.Printf("Ticker running at %v\n", time.Second/frameRate)
	defer ticker.Stop()
	timeout := time.NewTimer(time.Until(until))
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
	_ = f.Dev.Clear()
}
