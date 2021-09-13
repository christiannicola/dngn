package sdl

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/christiannicola/dngn/pkg/primitives"
	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	counter   int
	screen    *Screen
	IsRunning bool
}

func NewGame(width, height, fps int32, title string) (*Game, error) {
	var err error

	g := &Game{0, nil, true}

	if g.screen, err = NewScreen(width, height, title); err != nil {
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

	// NOTE (c.nicola): The game is running at 25 ticks a second
	if g.counter%(60/25) != 0 {
		return nil
	}

	for y := 0; y < g.screen.Terminal.Height(); y++ {
		for x := 0; x < g.screen.Terminal.Width(); x++ {
			if err := g.screen.Terminal.WriteChar(x, y, graphics.MediumShade, graphics.DarkGray, graphics.DarkRed); err != nil {
				panic(err)
			}
		}
	}

	box := primitives.NewRect(0, 0, g.screen.Terminal.Width()/2, g.screen.Terminal.Height()/2)

	g.screen.Terminal.Fill(box.X(), box.Y(), box.Width(), box.Height(), graphics.DarkBrown)

	t := "dngn\nis great"
	text := graphics.NewText(t, graphics.LightGray, graphics.Black)

	text.SetPos(graphics.TopLeft, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.TopCentered, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.TopRight, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.LeftCentered, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.Centered, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.RightCentered, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.BottomLeft, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.BottomCentered, box.X(), box.Y(), box.Width(), box.Height())
	if err := g.screen.Terminal.WriteText(text); err != nil {
		return err
	}

	text.SetPos(graphics.BottomRight, box.X(), box.Y(), box.Width(), box.Height())
	return g.screen.Terminal.WriteText(text)
}

func (g *Game) Draw() error {
	if err := g.screen.Terminal.Render(g.screen.DrawGlyph); err != nil {
		return err
	}

	g.screen.Renderer.Present()

	return nil
}

func (g *Game) Shutdown() error {
	return g.screen.Close()
}
