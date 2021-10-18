package graphics

import (
	"github.com/christiannicola/dngn/pkg/primitives"
	"image"
)

type Surface interface {
	Renderer() Renderer
	Clear(color Color)
	DrawRect(width, height int, color Color)
	DrawLine(x, y int, color Color)
	GetSize() (width, height int)
	GetDepth() int
	Pop()
	PopN(n int)
	PushColor(color Color)
	PushTranslation(x, y int)
	PushSkew(x, y float64)
	PushScale(x, y float64)
	PushBrightness(brightness float64)
	PushSaturation(saturation float64)
	Render(surface Surface)
	RenderSection(surface Surface, bound primitives.Rect)
	ReplacePixels(pixels []byte)
	Screenshot() *image.RGBA
}
