package graphics

import (
	"golang.org/x/image/font"
	"image"
	"image/color"
)

type Surface interface {
	Renderer() Renderer
	Clear(color color.Color)
	DrawRect(width, height int, color color.Color)
	DrawLine(x, y int, color color.Color)
	DrawText(x, y int, text string, face font.Face, clr color.Color)
	DrawGlyph(x, y int, glyph rune, face font.Face, clr color.Color)
	GetSize() (width, height int)
	GetDepth() int
	Pop()
	PopN(n int)
	PushColor(color color.Color)
	PushTranslation(x, y int)
	PushSkew(x, y float64)
	PushScale(x, y float64)
	PushBrightness(brightness float64)
	PushSaturation(saturation float64)
	Render(surface Surface)
	RenderSection(surface Surface, bound image.Rectangle)
	ReplacePixels(pixels []byte)
	Screenshot() *image.RGBA
}
