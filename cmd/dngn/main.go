package main

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

const (
	glyphWidth  = 16
	glyphHeight = 32
)

type SDLTerminal struct {
	terminal graphics.Terminal
	window   *sdl.Window
	renderer *sdl.Renderer
	surface  *sdl.Surface
	font     *ttf.Font
}

func NewSDLTerminal(width, height int32, title string) (*SDLTerminal, error) {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	if err = ttf.Init(); err != nil {
		return nil, err
	}

	terminal := SDLTerminal{}

	terminal.terminal = graphics.NewTerminal(int(width), int(height))

	if terminal.window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width*glyphWidth, height*glyphHeight, sdl.WINDOW_SHOWN); err != nil {
		return nil, err
	}

	if terminal.surface, err = terminal.window.GetSurface(); err != nil {
		return nil, err
	}

	if terminal.renderer, err = terminal.window.GetRenderer(); err != nil {
		return nil, err
	}

	if terminal.font, err = ttf.OpenFont("../assets/MxPlus_IBM_BIOS-2y.ttf", 16); err != nil {
		return nil, err
	}

	return &terminal, nil
}

func (s SDLTerminal) convertColor(c graphics.Color) sdl.Color {
	return sdl.Color{R: uint8(c.R()), G: uint8(c.G()), B: uint8(c.B()), A: 100}
}

func (s *SDLTerminal) generateGlyphTexture(g graphics.Glyph) (*sdl.Texture, error) {
	surface, err := s.font.RenderGlyphShaded(g.Rune(), s.convertColor(g.ForeGroundColor()), s.convertColor(g.BackgroundColor()))
	if err != nil {
		return nil, err
	}

	defer surface.Free()

	return s.renderer.CreateTextureFromSurface(surface)
}

func (s *SDLTerminal) drawGlyph(x, y int, g graphics.Glyph) error {
	texture, err := s.generateGlyphTexture(g)
	defer texture.Destroy()

	if err != nil {
		return err
	}

	return s.renderer.Copy(texture, nil, &sdl.Rect{
		X: int32(x * glyphWidth),
		Y: int32(y * glyphHeight),
		W: glyphWidth,
		H: glyphHeight,
	})
}

func (s *SDLTerminal) close() error {
	s.font.Close()

	if err := s.window.Destroy(); err != nil {
		return err
	}

	sdl.Quit()

	return nil
}

func main() {
	const width, height = 100, 40
	const fps = 60
	const frameDelay = 1000 / fps

	terminal, err := NewSDLTerminal(width, height, "dngn")

	if err != nil {
		panic(err)
	}

	isRunning := true

	var frameStart, frameTime uint32 = 1, 1

	for isRunning {

		frameStart = sdl.GetTicks()

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				event := sdl.PollEvent()
				switch event.(type) {
				case *sdl.QuitEvent:
					isRunning = false
				}

				var fg, bg graphics.Color

				switch time.Now().Nanosecond() % 10 {
				case 0:
					fallthrough
				case 9:
					fg = graphics.Red
					bg = graphics.DarkRed
				case 1:
					fallthrough
				case 8:
					fg = graphics.Blue
					bg = graphics.DarkBlue
				case 2:
					fallthrough
				case 7:
					fg = graphics.Yellow
					bg = graphics.DarkYellow
				case 3:
					fallthrough
				case 6:
					fg = graphics.Gold
					bg = graphics.DarkGold
				case 4:
					fallthrough
				case 5:
					fg = graphics.Green
					bg = graphics.DarkGreen
				}

				if err = terminal.terminal.WriteChar(x, y, graphics.MediumShade, fg, bg); err != nil {
					panic(err)
				}
			}
		}

		if err = terminal.terminal.Render(terminal.drawGlyph); err != nil {
			panic(err)
		}

		terminal.renderer.Present()

		frameTime = sdl.GetTicks() - frameStart

		if frameDelay > frameTime {
			sdl.Delay(frameDelay - frameTime)
		}
	}

	if err = terminal.close(); err != nil {
		panic(err)
	}
}
