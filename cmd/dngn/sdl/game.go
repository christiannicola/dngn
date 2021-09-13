package sdl

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/christiannicola/dngn/pkg/ui"
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
	var event sdl.Event
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			g.IsRunning = false
		}
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
			if err := g.screen.Terminal.WriteChar(x, y, graphics.MediumShade, graphics.DarkerWarmGray, graphics.Black); err != nil {
				panic(err)
			}
		}
	}

	logo := graphics.NewText(ui.Logo, graphics.Buttermilk, graphics.DarkWarmGray)
	start := graphics.NewText("   New Game   ", graphics.Buttermilk, graphics.DarkWarmGray)
	load := graphics.NewText("   Load Game  ", graphics.Buttermilk, graphics.DarkWarmGray)
	quit := graphics.NewText("     Quit     ", graphics.Buttermilk, graphics.DarkWarmGray)
	copyright := graphics.NewText("2021 c Christian Nicola", graphics.Buttermilk, graphics.DarkWarmGray)

	logo.SetPos(graphics.TopCentered, 0, 3, g.screen.Terminal.Width(), g.screen.Terminal.Height())
	start.SetPos(graphics.Centered, 0, 0, g.screen.Terminal.Width(), g.screen.Terminal.Height())
	load.SetPos(graphics.Centered, 0, 3, g.screen.Terminal.Width(), g.screen.Terminal.Height())
	quit.SetPos(graphics.Centered, 0, 6, g.screen.Terminal.Width(), g.screen.Terminal.Height())
	copyright.SetPos(graphics.BottomLeft, 0, 0, g.screen.Terminal.Width(), g.screen.Terminal.Height())

	if err := g.screen.Terminal.WriteText(start); err != nil {
		return err
	}
	if err := g.screen.Terminal.WriteText(load); err != nil {
		return err
	}
	if err := g.screen.Terminal.WriteText(quit); err != nil {
		return err
	}

	if err := g.screen.Terminal.WriteText(copyright); err != nil {
		return err
	}
	return g.screen.Terminal.WriteText(logo)
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
