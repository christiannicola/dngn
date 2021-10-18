package ebiten

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/christiannicola/dngn/pkg/graphics/ebiten"
	"github.com/christiannicola/dngn/pkg/ui"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
	"os"
)

const glyphWidth = 16

type Game struct {
	IsRunning bool
	Terminal  *graphics.Terminal
	Renderer  graphics.Renderer
	title     string
}

func NewGame(width, height int, title string) (*Game, error) {
	var err error

	game := Game{IsRunning: true, title: title}

	fontFile, err := os.OpenFile("assets/MxPlus_IBM_BIOS-2y.ttf", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer fontFile.Close()

	fontContent, err := ioutil.ReadAll(fontFile)
	if err != nil {
		return nil, err
	}

	parsedFont, err := opentype.Parse(fontContent)
	if err != nil {
		return nil, err
	}

	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{Size: glyphWidth, DPI: 72, Hinting: font.HintingFull})
	if err != nil {
		return nil, err
	}

	game.Terminal = graphics.NewTerminal(width, height, face)

	game.Renderer, err = ebiten.CreateRenderer()
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (g *Game) Update() error {
	for y := 0; y < g.Terminal.Height(); y++ {
		for x := 0; x < g.Terminal.Width(); x++ {
			if err := g.Terminal.WriteChar(x, y, graphics.MediumShade, graphics.DarkerWarmGray, graphics.Black); err != nil {
				panic(err)
			}
		}
	}

	logo := graphics.NewText(ui.Logo, graphics.Buttermilk, graphics.DarkWarmGray)
	start := graphics.NewText("   New Game   ", graphics.Buttermilk, graphics.DarkWarmGray)
	load := graphics.NewText("   Load Game  ", graphics.Buttermilk, graphics.DarkWarmGray)
	quit := graphics.NewText("     Quit     ", graphics.Buttermilk, graphics.DarkWarmGray)
	copyright := graphics.NewText("2021 c Christian Nicola", graphics.Buttermilk, graphics.DarkWarmGray)

	logo.SetPos(graphics.TopCentered, 0, 3, g.Terminal.Width(), g.Terminal.Height())
	start.SetPos(graphics.Centered, 0, 0, g.Terminal.Width(), g.Terminal.Height())
	load.SetPos(graphics.Centered, 0, 3, g.Terminal.Width(), g.Terminal.Height())
	quit.SetPos(graphics.Centered, 0, 6, g.Terminal.Width(), g.Terminal.Height())
	copyright.SetPos(graphics.BottomLeft, 0, 0, g.Terminal.Width(), g.Terminal.Height())

	if err := g.Terminal.WriteText(start); err != nil {
		return err
	}
	if err := g.Terminal.WriteText(load); err != nil {
		return err
	}
	if err := g.Terminal.WriteText(quit); err != nil {
		return err
	}

	if err := g.Terminal.WriteText(copyright); err != nil {
		return err
	}
	return g.Terminal.WriteText(logo)
}

func (g *Game) Draw(surface graphics.Surface) error {
	return g.Terminal.Render(surface)
}
