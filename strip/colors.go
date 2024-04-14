package strip

type RGBColor struct {
	R, G, B uint8
}

// LED represents an RGB color where R, G, and B are the red, green, and blue values respectively.
type Color struct {
	R, G, B int16
}

// LEDStrip is a slice of LEDs, representing the entire strip of LEDs.
type LEDStrip []Color

func Color2RGB(c Color) RGBColor {
	r := to8(c.R)
	g := to8(c.G)
	b := to8(c.B)
	return RGBColor{R: r, G: g, B: b}
}

func to8(c int16) uint8 {
	var ret uint8
	if c > 255 {
		ret = 255
	} else if c < 0 {
		ret = 0
	} else {
		ret = uint8(c)
	}
	return ret
}

// windToRotation converts a wind speed value (m/s) to a rotation value.
// Rotation is 0 when wind speed is 0.
// If a tiny bit of wind, rotation is 10, which means the colors are shifted every 10 tick.
// If a lot of wind, rotation is 3, which means the colors are shifted every 3 ticks.
// It maxes out at 1, which means the colors are shifted every tick.
func windToRotation(speed float64) int {
	if speed >= 8.5 {
		return 1
	}
	if speed >= 6.5 {
		return 2
	}
	if speed >= 4.5 {
		return 4
	}
	if speed >= 2.5 {
		return 8
	}
	if speed >= 1.5 {
		return 16
	}
	if speed >= 0.5 {
		return 32
	}
	if speed >= 0.1 {
		return 64
	}
	return 0
}
