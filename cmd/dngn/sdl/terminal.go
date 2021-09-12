package sdl

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	glyphWidth  = 16
	glyphHeight = 32
)

type Terminal struct {
	Internal *graphics.Terminal
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Surface  *sdl.Surface
	Font     *ttf.Font
}

func NewTerminal(width, height int32, title string) (*Terminal, error) {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if err = ttf.Init(); err != nil {
		return nil, err
	}

	terminal := Terminal{}

	terminal.Internal = graphics.NewTerminal(int(width), int(height))

	if terminal.Window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*glyphWidth, height*glyphHeight, sdl.WINDOW_SHOWN); err != nil {
		return nil, err
	}

	if terminal.Surface, err = terminal.Window.GetSurface(); err != nil {
		return nil, err
	}

	if terminal.Renderer, err = terminal.Window.GetRenderer(); err != nil {
		return nil, err
	}

	if terminal.Font, err = ttf.OpenFont("../assets/MxPlus_IBM_BIOS-2y.ttf", 16); err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (s Terminal) convertColor(c graphics.Color) sdl.Color {
	return sdl.Color{R: uint8(c.R()), G: uint8(c.G()), B: uint8(c.B()), A: 100}
}

func (s *Terminal) generateGlyphTexture(g graphics.Glyph) (*sdl.Texture, error) {
	surface, err := s.Font.RenderGlyphShaded(g.Rune(), s.convertColor(g.ForeGroundColor()), s.convertColor(g.BackgroundColor()))
	if err != nil {
		return nil, err
	}

	defer surface.Free()

	return s.Renderer.CreateTextureFromSurface(surface)
}

func (s *Terminal) DrawGlyph(x, y int, g graphics.Glyph) error {
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

func (s *Terminal) Close() error {
	s.Font.Close()

	if err := s.Window.Destroy(); err != nil {
		return err
	}

	sdl.Quit()

	return nil
}
