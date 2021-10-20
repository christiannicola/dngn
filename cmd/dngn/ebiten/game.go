package ebiten

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/christiannicola/dngn/pkg/graphics/ebiten"
)

type Game struct {
	IsRunning bool
	Renderer  graphics.Renderer
	title     string
}

func NewGame(title string) (*Game, error) {
	var err error

	game := Game{IsRunning: true, title: title}

	game.Renderer, err = ebiten.CreateRenderer()
	if err != nil {
		return nil, err
	}

	return &game, nil
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(surface graphics.Surface) error {
	target := g.Renderer.NewSurface(50, 50)
	target.DrawRect(50, 50, graphics.DarkerCoolGray)

	surface.Render(target)

	return nil
}
