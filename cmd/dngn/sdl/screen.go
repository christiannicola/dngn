package sdl

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	glyphWidth  = 16
	glyphHeight = 16
)

type Screen struct {
	Terminal *graphics.Terminal
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Font     *ttf.Font
}

func NewScreen(width, height int32, title string) (*Screen, error) {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if err = ttf.Init(); err != nil {
		return nil, err
	}

	screen := Screen{}

	screen.Terminal = graphics.NewTerminal(int(width), int(height))

	if screen.Window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*glyphWidth, height*glyphHeight, sdl.WINDOW_SHOWN); err != nil {
		return nil, err
	}

	if screen.Renderer, err = sdl.CreateRenderer(screen.Window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return nil, err
	}

	if screen.Font, err = ttf.OpenFont("../assets/MxPlus_IBM_BIOS-2y.ttf", 16); err != nil {
		return nil, err
	}

	return &screen, nil
}

func (s *Screen) generateGlyphTexture(g graphics.Glyph) (*sdl.Texture, error) {
	surface, err := s.Font.RenderGlyphShaded(g.Rune(), sdl.Color(g.ForeGroundColor().RGBA()), sdl.Color(g.BackgroundColor().RGBA()))
	if err != nil {
		return nil, err
	}

	defer surface.Free()

	return s.Renderer.CreateTextureFromSurface(surface)
}

func (s *Screen) DrawGlyph(x, y int, g graphics.Glyph) error {
	texture, err := s.generateGlyphTexture(g)
	defer texture.Destroy()

	if err != nil {
		return err
	}

	return s.Renderer.Copy(texture, nil, &sdl.Rect{
		X: int32(x * glyphWidth),
		Y: int32(y * glyphHeight),
		W: glyphWidth,
		H: glyphHeight,
	})
}

func (s *Screen) Close() error {
	s.Font.Close()

	if err := s.Window.Destroy(); err != nil {
		return err
	}

	sdl.Quit()

	return nil
}
