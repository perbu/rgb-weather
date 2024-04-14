package strip

import (
	"math/rand"
)

// Hour represents a single hours of weather data.
type Hour struct {
	Forecast       *Forecast
	start          int
	rotationOffset int
	Precipitation  float64
	CloudCover     float64
	WindSpeed      float64
	leds           []Color
	noise          []Color
	rainDrops      []bool
}

func (f *Forecast) NewHour(offset, leds int) *Hour {
	const noiseAmplitude = 25
	noise := make([]Color, leds)
	for i := range noise {
		noise[i] = Color{
			R: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
			G: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
			B: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
		}
	}
	return &Hour{
		Forecast:  f,
		leds:      make([]Color, leds),
		noise:     noise,
		start:     offset * leds,
		rainDrops: make([]bool, leds),
	}
}

func (h *Hour) Update() {
	const rainUpdateInterval = 200
	// cloud cover:
	for i := 0; i < len(h.leds); i++ {
		h.leds[i] = CloudCoverToColor(h.CloudCover)
	}
	rotationSpeed := windToRotation(h.WindSpeed)
	if rotationSpeed > 0 {
		if h.Forecast.Ticks%rotationSpeed == 0 {
			h.rotationOffset = (h.rotationOffset + 1) % len(h.leds)
		}
	}
	ledsCopy := make([]Color, len(h.leds))

	// Apply the noise and make a copy of the leds
	for i := 0; i < len(ledsCopy); i++ {
		ledsCopy[i] = applyNoise(h.leds[i], h.noise[i])
	}
	// raindrops:
	rainFactor := RainFactor(h.Precipitation)
	if h.Forecast.Ticks%rainUpdateInterval == 0 {
		for i := 0; i < len(h.rainDrops); i++ {
			h.rainDrops[i] = rand.Float64() < rainFactor
		}
	}
	// overwrite the leds with raindrops where needed:
	for i := 0; i < len(h.leds); i++ {
		if h.rainDrops[i] {
			ledsCopy[i] = Color{R: 0, G: 0, B: 255}
		}
	}

	// rotate the leds in the copy:
	ledsCopy = rotate(ledsCopy, h.rotationOffset)

	// draw the colors onto the strip
	for i := 0; i < len(h.leds); i++ {
		target := h.start + i
		h.Forecast.Strip[target] = Color2RGB(ledsCopy[i])
	}
}

func rotate[T any](slice []T, n int) []T {
	l := len(slice)
	if l == 0 {
		return slice // no rotation needed for empty or nil slices
	}
	n = n % l // handle rotations longer than the slice itself
	if n < 0 {
		n += l // convert negative rotations to positive rotations
	}
	if n == 0 {
		return slice // no rotation needed
	}
	return append(slice[l-n:], slice[:l-n]...)
}

func applyNoise(c Color, noise Color) Color {
	c.R += noise.R
	c.G += noise.G
	c.B += noise.B
	return c
}

// CloudCoverToColor converts a cloud cover value (0-1) to an RGB color.
// It is yellow when the cloud cover is 0, and white when the cloud cover is 1.
// noise is added to the color to make it more interesting and to let the rotation be more visible.
func CloudCoverToColor(cover float64) Color {
	r := int16(255 * cover)
	g := int16(255 * cover)
	b := int16(0)
	return Color{R: r, G: g, B: b}
}

// RainFactor returns a factor between 0 and X that represents the amount of rain.
// X is the maximum amount of rain that can be displayed, might be 0.75 for example.
func RainFactor(rain float64) float64 {
	const maxRain = 0.75
	if rain > 5 {
		return maxRain
	}
	if rain > 4 {
		return maxRain * 0.8
	}
	if rain > 3 {
		return maxRain * 0.6
	}
	if rain > 2 {
		return maxRain * 0.4
	}
	if rain > 0 {
		return maxRain * 0.2
	}
	return 0
}
