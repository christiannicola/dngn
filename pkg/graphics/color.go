package graphics

import (
	"github.com/christiannicola/dngn/pkg/math"
	"image/color"
)

var (
	Black = Color{0, 0, 0}
	White = Color{255, 255, 255}

	LightGray = Color{192, 192, 192}
	Gray      = Color{128, 128, 128}
	DarkGray  = Color{64, 64, 64}

	LightRed = Color{255, 160, 160}
	Red      = Color{220, 0, 0}
	DarkRed  = Color{100, 0, 0}

	LightOrange = Color{255, 200, 170}
	Orange      = Color{255, 128, 0}
	DarkOrange  = Color{128, 64, 0}

	LightGold = Color{255, 230, 150}
	Gold      = Color{255, 192, 0}
	DarkGold  = Color{128, 96, 0}

	LightYellow = Color{255, 255, 150}
	Yellow      = Color{255, 255, 0}
	DarkYellow  = Color{128, 128, 0}

	LightGreen = Color{130, 255, 90}
	Green      = Color{0, 128, 0}
	DarkGreen  = Color{0, 64, 0}

	LightAqua = Color{128, 255, 255}
	Aqua      = Color{0, 255, 255}
	DarkAqua  = Color{0, 128, 128}

	LightBlue = Color{128, 160, 255}
	Blue      = Color{0, 64, 255}
	DarkBlue  = Color{0, 37, 168}

	LightPurple = Color{200, 140, 255}
	Purple      = Color{128, 0, 255}
	DarkPurple  = Color{64, 0, 128}

	LightBrown = Color{190, 150, 100}
	Brown      = Color{160, 110, 60}
	DarkBrown  = Color{100, 64, 32}

	Ash            = Color{0xe2, 0xdf, 0xf0}
	LightCoolGray  = Color{0x74, 0x92, 0xb5}
	CoolGray       = Color{0x3f, 0x4b, 0x73}
	DarkCoolGray   = Color{0x26, 0x2a, 0x42}
	DarkerCoolGray = Color{0x14, 0x13, 0x1f}

	LightWarmGray  = Color{0x84, 0x7e, 0x87}
	WarmGray       = Color{0x48, 0x40, 0x4a}
	DarkWarmGray   = Color{0x2a, 0x24, 0x2b}
	DarkerWarmGray = Color{0x16, 0x11, 0x17}

	Sandal = Color{0xbd, 0x90, 0x6c}
	Tan    = Color{0x8e, 0x52, 0x37}

	Carrot    = Color{0xb3, 0x4a, 0x04}
	Persimmon = Color{0x6e, 0x20, 0x0d}

	Buttermilk = Color{0xff, 0xee, 0xa8}
	Olive      = Color{0x63, 0x57, 0x07}
	DarkOlive  = Color{0x33, 0x30, 0x1c}

	Mint     = Color{0x81, 0xd9, 0x75}
	Lima     = Color{0x83, 0x9e, 0x0d}
	PeaGreen = Color{0x16, 0x75, 0x26}
	Sherwood = Color{0x00, 0x40, 0x27}

	Lavender = Color{0xc9, 0xa6, 0xff}
	Lilac    = Color{0xad, 0x58, 0xdb}
	Violet   = Color{0x38, 0x10, 0x7d}

	Pink   = Color{0xff, 0x7a, 0x69}
	Maroon = Color{0x54, 0x00, 0x27}
)

type Color struct {
	r, g, b int
}

func NewColor(r, g, b int) Color {
	return Color{r, g, b}
}

func (c Color) Add(other Color, fraction float32) Color {
	return Color{
		math.Clamp(int(float32(c.r)+float32(other.r)*fraction), 0, 255),
		math.Clamp(int(float32(c.g)+float32(other.g)*fraction), 0, 255),
		math.Clamp(int(float32(c.b)+float32(other.b)*fraction), 0, 255),
	}
}

func (c Color) Blend(other Color, otherFraction float32) Color {
	thisFraction := 1.0 - otherFraction

	return Color{
		math.Clamp(int(float32(c.r)*thisFraction+float32(other.r)*otherFraction), 0, 255),
		math.Clamp(int(float32(c.g)*thisFraction+float32(other.g)*otherFraction), 0, 255),
		math.Clamp(int(float32(c.b)*thisFraction+float32(other.b)*otherFraction), 0, 255),
	}
}

func (c Color) BlendPercent(other Color, percent int) Color {
	return c.Blend(other, float32(percent/100))
}

func (c Color) Equal(other Color) bool {
	return c.r == other.r && c.g == other.g && c.b == other.b
}

func (c Color) R() int {
	return c.r
}

func (c Color) G() int {
	return c.g
}

func (c Color) B() int {
	return c.b
}

func (c Color) RGBA() color.RGBA {
	return color.RGBA{R: uint8(c.R()), G: uint8(c.G()), B: uint8(c.B()), A: 0}
}
