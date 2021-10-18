package ebiten

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/christiannicola/dngn/pkg/primitives"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"math"
)

var _ graphics.Surface = &Surface{}

const (
	maxAlpha       = 0xff
	cacheLimit     = 512
	transparency25 = 0.25
	transparency50 = 0.50
	transparency75 = 0.75
)

type (
	colorMCacheKey uint32

	colorMCacheEntry struct {
		colorMatrix ebiten.ColorM
		atime       int64
	}

	surfaceState struct {
		x, y           int
		filter         ebiten.Filter
		color          color.Color
		brightness     float64
		saturation     float64
		skewX, skewY   float64
		scaleX, scaleY float64
	}

	Surface struct {
		renderer       *Renderer
		stateStack     []surfaceState
		stateCurrent   surfaceState
		image          *ebiten.Image
		colorMCache    map[colorMCacheKey]*colorMCacheEntry
		monotonicClock int64
	}
)

func (s Surface) Renderer() graphics.Renderer {
	return s.renderer
}

func (s Surface) Clear(color graphics.Color) {
	s.image.Fill(color)
}

func (s Surface) DrawRect(width, height int, color graphics.Color) {
	ebitenutil.DrawRect(s.image, float64(s.stateCurrent.x), float64(s.stateCurrent.y), float64(width), float64(height), color)
}

func (s Surface) DrawLine(x, y int, color graphics.Color) {
	ebitenutil.DrawLine(s.image, float64(s.stateCurrent.x), float64(s.stateCurrent.y), float64(s.stateCurrent.x+x), float64(s.stateCurrent.y+y), color)
}

func (s Surface) GetSize() (width, height int) {
	return s.image.Size()
}

func (s Surface) GetDepth() int {
	return len(s.stateStack)
}

func (s *Surface) Pop() {
	count := len(s.stateStack)

	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

func (s *Surface) PopN(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}

func (s *Surface) PushColor(color graphics.Color) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.color = color
}

func (s *Surface) PushTranslation(x, y int) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

func (s *Surface) PushSkew(x, y float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.skewX += x
	s.stateCurrent.skewY += y
}

func (s *Surface) PushScale(x, y float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.scaleX += x
	s.stateCurrent.scaleY += y
}

func (s *Surface) PushBrightness(brightness float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.brightness += brightness
}

func (s *Surface) PushSaturation(saturation float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.saturation += saturation
}

func (s Surface) Render(surface graphics.Surface) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 1 || s.stateCurrent.saturation != 1 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.image.DrawImage(surface.(*Surface).image, opts)
}

func (s Surface) RenderSection(surface graphics.Surface, bound primitives.Rect) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.image.DrawImage(surface.(*Surface).image.SubImage(image.Rect(bound.TopLeft().X(), bound.TopLeft().Y(), bound.BottomRight().X(), bound.BottomRight().Y())).(*ebiten.Image), opts)
}

func (s Surface) ReplacePixels(pixels []byte) {
	s.image.ReplacePixels(pixels)
}

func (s Surface) Screenshot() *image.RGBA {
	width, height := s.GetSize()
	bounds := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: width, Y: height}}
	rgba := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba.Set(x, y, s.image.At(x, y))
		}
	}

	return rgba
}

func createSurface(r *Renderer, img *ebiten.Image, currentState ...surfaceState) *Surface {
	state := surfaceState{
		x:          0,
		y:          0,
		filter:     0,
		color:      graphics.Color{},
		brightness: defaultBrightness,
		saturation: defaultSaturation,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
	}

	if len(currentState) > 0 {
		state = currentState[0]
	}

	return &Surface{
		renderer:     r,
		image:        img,
		stateCurrent: state,
		colorMCache:  make(map[colorMCacheKey]*colorMCacheEntry),
	}
}

func (s *Surface) createDrawImageOptions() *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}

	if s.stateCurrent.skewX != 0 || s.stateCurrent.skewY != 0 {
		opts.GeoM.Skew(s.stateCurrent.skewX, s.stateCurrent.skewY)
	}

	if s.stateCurrent.scaleX != 1.0 || s.stateCurrent.scaleY != 1.0 {
		opts.GeoM.Scale(s.stateCurrent.scaleX, s.stateCurrent.scaleY)
	}

	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))

	opts.Filter = s.stateCurrent.filter

	if s.stateCurrent.color != nil {
		opts.ColorM = s.colorToColorM(s.stateCurrent.color)
	}

	return opts
}

func (s *Surface) now() int64 {
	s.monotonicClock++
	return s.monotonicClock
}

func (s *Surface) colorToColorM(clr color.Color) ebiten.ColorM {
	// RGBA() is in [0 - 0xffff]. Adjust them in [0 - 0xff].
	cr, cg, cb, ca := clr.RGBA()
	cr >>= 8
	cg >>= 8
	cb >>= 8
	ca >>= 8

	if ca == 0 {
		emptyColorM := ebiten.ColorM{}
		emptyColorM.Scale(0, 0, 0, 0)

		return emptyColorM
	}

	key := colorMCacheKey(cr | (cg << 8) | (cb << 16) | (ca << 24))
	e, ok := s.colorMCache[key]

	if ok {
		e.atime = s.now()
		return e.colorMatrix
	}

	if len(s.colorMCache) > cacheLimit {
		oldest := int64(math.MaxInt64)
		oldestKey := colorMCacheKey(0)

		for key, c := range s.colorMCache {
			if c.atime < oldest {
				oldestKey = key
				oldest = c.atime
			}
		}

		delete(s.colorMCache, oldestKey)
	}

	cm := ebiten.ColorM{}
	rf := float64(cr) / float64(ca)
	gf := float64(cg) / float64(ca)
	bf := float64(cb) / float64(ca)
	af := float64(ca) / maxAlpha
	cm.Scale(rf, gf, bf, af)

	e = &colorMCacheEntry{
		colorMatrix: cm,
		atime:       s.now(),
	}

	s.colorMCache[key] = e

	return e.colorMatrix
}
