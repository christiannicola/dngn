package graphics

import "github.com/christiannicola/dngn/pkg/math"

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
