package sdl

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"time"
)

type Game struct {
	counter int
	terminal                               *Terminal
	IsRunning                              bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewGame(width, height, fps int32, title string) (*Game, error) {
	var err error

	g := &Game{0,nil, true}

	if g.terminal, err = NewTerminal(width, height, title); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Game) HandleEvents() error {
	event := sdl.PollEvent()

	switch event.(type) {
	case *sdl.QuitEvent:
		g.IsRunning = false
	}

	return nil
}

func (g *Game) Update() error {
	g.counter++

	// NOTE (c.nicola): The game is running at 60 ticks a second, and we only update once every second
	if g.counter%60 != 0 {
		return nil
	}

	for y := 0; y < g.terminal.Internal.Height(); y++ {
		for x := 0; x < g.terminal.Internal.Width(); x++ {
			var fg, bg graphics.Color

			switch 0 /*rand.Intn(0x7f) % 10*/ {
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

			if err := g.terminal.Internal.WriteChar(x, y, graphics.MediumShade, fg, bg); err != nil {
				panic(err)
			}
		}
	}

	return nil
}

func (g *Game) Draw() error {
	if err := g.terminal.Internal.Render(g.terminal.DrawGlyph); err != nil {
		return err
	}

	g.terminal.Renderer.Present()

	return nil
}

func (g *Game) Shutdown() error {
	return g.terminal.Close()
}
