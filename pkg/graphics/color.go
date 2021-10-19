package graphics

import (
	"github.com/christiannicola/dngn/pkg/math"
	"image/color"
)

var (
	Black = color.RGBA{A: 0xff}
	White = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

	LightGray = color.RGBA{R: 0xc0, G: 0xc0, B: 0xc0, A: 0xff}
	Gray      = color.RGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xff}
	DarkGray  = color.RGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xff}

	LightRed = color.RGBA{R: 0xff, G: 0xa0, B: 0xa0, A: 0xff}
	Red      = color.RGBA{R: 0xdc, A: 0xff}
	DarkRed  = color.RGBA{R: 0x64, A: 0xff}

	LightOrange = color.RGBA{R: 0xff, G: 0xc8, B: 0xaa, A: 0xff}
	Orange      = color.RGBA{R: 0xff, G: 0x80, A: 0xff}
	DarkOrange  = color.RGBA{R: 0x80, G: 0x40, A: 0xff}

	LightGold = color.RGBA{R: 0xff, G: 0xe6, B: 0x96, A: 0xff}
	Gold      = color.RGBA{R: 0xff, G: 0xc0, A: 0xff}
	DarkGold  = color.RGBA{R: 0x80, G: 0x60, A: 0xff}

	LightYellow = color.RGBA{R: 0xff, G: 0xff, B: 0x96, A: 0xff}
	Yellow      = color.RGBA{R: 0xff, G: 0xff, A: 0xff}
	DarkYellow  = color.RGBA{R: 0x80, G: 0x80, A: 0xff}

	LightGreen = color.RGBA{R: 0x82, G: 0xff, B: 0x5a, A: 0xff}
	Green      = color.RGBA{G: 0x80, A: 0xff}
	DarkGreen  = color.RGBA{G: 0x40, A: 0xff}

	LightAqua = color.RGBA{R: 0x80, G: 0xff, B: 0xff, A: 0xff}
	Aqua      = color.RGBA{G: 0xff, B: 0xff, A: 0xff}
	DarkAqua  = color.RGBA{G: 0x80, B: 0x80, A: 0xff}

	LightBlue = color.RGBA{R: 0x80, G: 0xa0, B: 0xff, A: 0xff}
	Blue      = color.RGBA{G: 0x40, B: 0xff, A: 0xff}
	DarkBlue  = color.RGBA{G: 0x25, B: 0xa8, A: 0xff}

	LightPurple = color.RGBA{R: 0xc8, G: 0x8c, B: 0xff, A: 0xff}
	Purple      = color.RGBA{R: 0x80, B: 0xff, A: 0xff}
	DarkPurple  = color.RGBA{R: 0x40, B: 0x80, A: 0xff}

	LightBrown = color.RGBA{R: 0xbe, G: 0x96, B: 0x64, A: 0xff}
	Brown      = color.RGBA{R: 0xa0, G: 0x6e, B: 0x3c, A: 0xff}
	DarkBrown  = color.RGBA{R: 0x64, G: 0x40, B: 0x20, A: 0xff}

	Ash            = color.RGBA{R: 0xe2, G: 0xdf, B: 0xf0, A: 0xff}
	LightCoolGray  = color.RGBA{R: 0x74, G: 0x92, B: 0xb5, A: 0xff}
	CoolGray       = color.RGBA{R: 0x3f, G: 0x4b, B: 0x73, A: 0xff}
	DarkCoolGray   = color.RGBA{R: 0x26, G: 0x2a, B: 0x42, A: 0xff}
	DarkerCoolGray = color.RGBA{R: 0x14, G: 0x13, B: 0x1f, A: 0xff}

	LightWarmGray  = color.RGBA{R: 0x84, G: 0x7e, B: 0x87, A: 0xff}
	WarmGray       = color.RGBA{R: 0x48, G: 0x40, B: 0x4a, A: 0xff}
	DarkWarmGray   = color.RGBA{R: 0x2a, G: 0x24, B: 0x2b, A: 0xff}
	DarkerWarmGray = color.RGBA{R: 0x16, G: 0x11, B: 0x17, A: 0xff}

	Sandal = color.RGBA{R: 0xbd, G: 0x90, B: 0x6c, A: 0xff}
	Tan    = color.RGBA{R: 0x8e, G: 0x52, B: 0x37, A: 0xff}

	Carrot    = color.RGBA{R: 0xb3, G: 0x4a, B: 0x04, A: 0xff}
	Persimmon = color.RGBA{R: 0x6e, G: 0x20, B: 0x0d, A: 0xff}

	Buttermilk = color.RGBA{R: 0xff, G: 0xee, B: 0xa8, A: 0xff}
	Olive      = color.RGBA{R: 0x63, G: 0x57, B: 0x07, A: 0xff}
	DarkOlive  = color.RGBA{R: 0x33, G: 0x30, B: 0x1c, A: 0xff}

	Mint     = color.RGBA{R: 0x81, G: 0xd9, B: 0x75, A: 0xff}
	Lima     = color.RGBA{R: 0x83, G: 0x9e, B: 0x0d, A: 0xff}
	PeaGreen = color.RGBA{R: 0x16, G: 0x75, B: 0x26, A: 0xff}
	Sherwood = color.RGBA{G: 0x40, B: 0x27, A: 0xff}

	Lavender = color.RGBA{R: 0xc9, G: 0xa6, B: 0xff, A: 0xff}
	Lilac    = color.RGBA{R: 0xad, G: 0x58, B: 0xdb, A: 0xff}
	Violet   = color.RGBA{R: 0x38, G: 0x10, B: 0x7d, A: 0xff}

	Pink   = color.RGBA{R: 0xff, G: 0x7a, B: 0x69, A: 0xff}
	Maroon = color.RGBA{R: 0x54, B: 0x27, A: 0xff}
)

func NewColor(r, g, b uint8) color.Color {
	return color.RGBA{r, g, b, 0xff}
}

func AddColors(first, other color.Color, fraction float32) color.Color {
	r1, g1, b1, _ := first.RGBA()
	r2, g2, b2, _ := other.RGBA()

	return color.RGBA{
		R: math.ClampUint8(uint8(float32(r1)+float32(r2)*fraction), 0x00, 0xff),
		G: math.ClampUint8(uint8(float32(g1)+float32(g2)*fraction), 0x00, 0xff),
		B: math.ClampUint8(uint8(float32(b1)+float32(b2)*fraction), 0x00, 0xff),
		A: 0xff,
	}
}

func BlendColors(first, other color.Color, otherFraction float32) color.Color {
	r1, g1, b1, _ := first.RGBA()
	r2, g2, b2, _ := other.RGBA()
	thisFraction := 1.0 - otherFraction

	return color.RGBA{
		R: math.ClampUint8(uint8(float32(r1)*thisFraction+float32(r2)*otherFraction), 0x00, 0xff),
		G: math.ClampUint8(uint8(float32(g1)*thisFraction+float32(g2)*otherFraction), 0x00, 0xff),
		B: math.ClampUint8(uint8(float32(b1)*thisFraction+float32(b2)*otherFraction), 0x00, 0xff),
		A: 0xff,
	}
}

func BlendColorsPercent(first, other color.Color, percent int) color.Color {
	return BlendColors(first, other, float32(percent/100))
}

func AreColorsEqual(first, other color.Color) bool {
	r1, g1, b1, a1 := first.RGBA()
	r2, g2, b2, a2 := other.RGBA()

	return r1 == r2 && g1 == g2 && b1 == b2 && a1 == a2
}
