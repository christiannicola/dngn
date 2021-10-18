package ebiten

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
)

const (
	screenWidth       = 800
	screenHeight      = 600
	defaultSaturation = 1.0
	defaultBrightness = 1.0
	defaultSkewX      = 0.0
	defaultSkewY      = 0.0
	defaultScaleX     = 1.0
	defaultScaleY     = 1.0
)

var _ graphics.Renderer = &Renderer{}

type Renderer struct {
	graphics.UpdateCallback
	graphics.RenderCallback
	lastError error
}

func CreateRenderer() (*Renderer, error) {
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetFullscreen(false)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetMaxTPS(60)

	return &Renderer{}, nil
}

func (r *Renderer) Update() error {
	if r.UpdateCallback == nil {
		return errNoUpdateCallback
	}

	return r.UpdateCallback()
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.lastError = nil

	if r.RenderCallback == nil {
		r.lastError = errNoRenderCallback

		return
	}

	r.lastError = r.RenderCallback(createSurface(r, screen))
}

func (r *Renderer) Layout(_, _ int) (width, height int) {
	return screenWidth, screenHeight
}

func (r *Renderer) GetRendererName() string {
	return "ebiten"
}

func (r *Renderer) SetWindowIcon(path string) {
	_, icon, err := ebitenutil.NewImageFromFile(path)

	if err == nil {
		ebiten.SetWindowIcon([]image.Image{icon})
	}
}

func (r *Renderer) Run(rc graphics.RenderCallback, u graphics.UpdateCallback, width, height int, title string) error {
	r.RenderCallback = rc
	r.UpdateCallback = u

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(width, height)

	return ebiten.RunGame(r)
}

func (r *Renderer) IsDrawingSkipped() bool {
	return r.lastError != nil
}

func (r *Renderer) CreateSurface(surface graphics.Surface) (graphics.Surface, error) {
	img := surface.(*Surface).image
	sfcState := surfaceState{
		filter:     ebiten.FilterNearest,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
	}
	result := createSurface(r, img, sfcState)

	return result, nil
}

func (r *Renderer) NewSurface(width, height int) graphics.Surface {
	img := ebiten.NewImage(width, height)

	return createSurface(r, img)
}

func (r *Renderer) IsFullScreen() bool {
	return ebiten.IsFullscreen()
}

func (r *Renderer) SetFullScreen(fullScreen bool) {
	ebiten.SetFullscreen(fullScreen)
}

func (r *Renderer) SetVSyncEnabled(vsync bool) {
	ebiten.SetVsyncEnabled(vsync)
}

func (r *Renderer) IsVSyncEnabled() bool {
	return ebiten.IsVsyncEnabled()
}

func (r *Renderer) GetCursorPos() (int, int) {
	return ebiten.CursorPosition()
}

func (r *Renderer) CurrentFPS() float64 {
	return ebiten.CurrentFPS()
}
