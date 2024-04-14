package strip

import (
	"math/rand"
)

// Hour represents a single hours of weather data.
type Hour struct {
	Forecast       *Forecast
	start          int
	rotationOffset int
	// raindrops      []int
	// precipitation float64
	cloudCover float64
	windSpeed  float64
	leds       []Color
	noise      []Color
}

func (f *Forecast) NewHour(offset, leds int) *Hour {
	const noiseAmplitude = 5
	noise := make([]Color, leds)
	for i := range noise {
		noise[i] = Color{
			R: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
			G: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
			B: int16(rand.Intn(noiseAmplitude*2) - noiseAmplitude),
		}
	}
	return &Hour{
		Forecast: f,
		leds:     make([]Color, leds),
		noise:    noise,
		start:    offset * leds,
	}
}

func (h *Hour) Update() {
	// cloud cover:
	for i := 0; i < len(h.leds); i++ {
		h.leds[i] = h.CloudCoverToColor(h.cloudCover)
	}
	rotationSpeed := windToRotation(h.windSpeed)
	if rotationSpeed > 0 {
		if h.Forecast.Ticks%rotationSpeed == 0 {
			h.rotationOffset = (h.rotationOffset + 1) % len(h.leds)
		}
	}
	ledsCopy := make([]Color, len(h.leds))
	copy(ledsCopy, h.leds)
	for i := 0; i < len(ledsCopy); i++ {
		ledsCopy[i] = applyNoise(ledsCopy[i], h.noise[i])
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
	c.R = +noise.R
	c.G = +noise.G
	c.B = +noise.B
	return c
}

// CloudCoverToColor converts a cloud cover value (0-1) to an RGB color.
// It is yellow when the cloud cover is 0, and white when the cloud cover is 1.
// noise is added to the color to make it more interesting and to let the rotation be more visible.
func (h *Hour) CloudCoverToColor(cover float64) Color {
	r := int16(255 * cover)
	g := int16(255 * cover)
	b := int16(255)
	return Color{R: r, G: g, B: b}
}
