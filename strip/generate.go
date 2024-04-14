package strip

import "math/rand"

func GenerateForecast(hours, ledsPerHour int) *Forecast {
	leds := hours * ledsPerHour
	f := &Forecast{
		Strip: make([]RGBColor, leds),
	}
	f.Hours = make([]*Hour, hours)
	for i := range f.Hours {
		f.Hours[i] = f.NewHour(i, ledsPerHour)
		f.Hours[i].CloudCover = rand.Float64()
		f.Hours[i].WindSpeed = rand.Float64() * 10
		f.Hours[i].Precipitation = rand.Float64() * 4
	}
	return f
}
