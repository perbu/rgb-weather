package strip

import (
	"fmt"
	"math/rand"
)

// Hour represents a single hours of weather data.
type Hour struct {
	len           int
	start         int
	Precipitation float64
	CloudCover    float64
	WindSpeed     float64
	rainDrops     []bool
	forecast      *Forecast
}

func (f *Forecast) NewHour(offset, leds int) *Hour {
	return &Hour{
		len: leds,
		//noise:     noise,
		start:     offset * leds,
		rainDrops: make([]bool, leds),
		forecast:  f,
	}
}

func (h *Hour) Update(ticks int) {
	const rainUpdateInterval = 200 // update raindrops every 200 ticks
	// cloud cover:
	leds := make([]Color, h.len)
	// skip the first led to make the rotation more visible
	for i := 1; i < h.len; i++ {
		leds[i] = CloudCoverToColor(h.CloudCover / 100)
	}
	rotationSpeed := windToRotation(h.WindSpeed) // the higher the speed the slower the rotation.
	var rotationOffset int
	if rotationSpeed > 0 {
		rotationOffset = int(float64(ticks) / float64(rotationSpeed))
	}
	fmt.Println("rotationOffset:", rotationOffset, "rotationSpeed:", rotationSpeed)
	// raindrops:
	rainFactor := RainFactor(h.Precipitation)
	if ticks%rainUpdateInterval == 0 {
		for i := 0; i < len(h.rainDrops); i++ {
			h.rainDrops[i] = rand.Float64() < rainFactor
		}
	}
	// overwrite the leds with raindrops where needed:
	for i := 0; i < h.len; i++ {
		if h.rainDrops[i] {
			leds[i] = Color{R: 0, G: 0, B: 255}
		}
	}
	// ledsCopy[0] = Color{R: 0, G: 0, B: 0}
	leds = rotate(leds, rotationOffset)

	// draw the colors onto the strip
	for i := 0; i < h.len; i++ {
		target := h.start + i
		h.forecast.Strip[target] = Color2RGB(leds[i])
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
// getColorFromValue returns the RGB values based on the input value
func CloudCoverToColor(value float64) Color {
	if value < 0.0 || value > 1.0 {
		fmt.Println("Value should be between 0.0 and 1.0, was: ", value)
		return Color{}
	}

	// Define start and end colors
	rStart, gStart, bStart := uint8(255), uint8(255), uint8(0) // Bright yellow
	rEnd, gEnd, bEnd := uint8(32), uint8(32), uint8(32)        // Dark gray

	// Interpolate each color channel
	r := interpolate(rStart, rEnd, value)
	g := interpolate(gStart, gEnd, value)
	b := interpolate(bStart, bEnd, value)

	return Color{R: int16(r), G: int16(g), B: int16(b)}
}

// interpolate calculates the linear interpolation for a single channel
func interpolate(start, end uint8, value float64) uint8 {
	return uint8(float64(start) + (float64(end)-float64(start))*value)
}

// RainFactor returns a factor between 0 and X that represents the amount of rain.
// X is the maximum amount of rain that can be displayed, might be 0.75 for example.
func RainFactor(rain float64) float64 {
	const maxRain = 0.80
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
