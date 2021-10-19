package graphics

import (
	"github.com/christiannicola/dngn/pkg/primitives"
	"golang.org/x/image/font"
	"image/color"
)

type (
	Terminal struct {
		display Display
		bg, fg  color.Color
		face    font.Face
	}
)

func NewTerminal(width, height int, face font.Face) *Terminal {
	return &Terminal{NewDisplay(width, height), Black, White, face}
}

func (t Terminal) Width() int {
	return t.display.Width()
}

func (t Terminal) Height() int {
	return t.display.Height()
}

func (t Terminal) Size() primitives.Vector {
	return t.display.Size()
}

func (t *Terminal) Fill(x, y, width, height int, color color.Color) error {
	glyph := NewGlyphFromCharCode(Space, t.fg, color)

	for iy := y; iy < y+height; iy++ {
		for ix := x; ix < x+width; ix++ {
			if err := t.display.SetGlyph(ix, iy, glyph); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Terminal) WriteChar(x, y int, c CharCode, fg, bg color.Color) error {
	return t.display.SetGlyph(x, y, Glyph{c, fg, bg})
}

func (t *Terminal) WriteString(x, y int, text string, fg, bg color.Color) error {
	runes := []rune(text)

	for i := range runes {
		if x+i >= t.Width() {
			break
		}

		if err := t.display.SetGlyph(x+i, y, NewGlyphFromRune(runes[i], fg, bg)); err != nil {
			return err
		}
	}

	return nil
}

func (t *Terminal) WriteText(text *Text) error {
	for iy := text.y; iy < text.Size().Y()+text.y; iy++ {
		for ix := text.x; ix < text.Size().X()+text.x; ix++ {
			gl, err := text.glyphs.Get(ix-text.x, iy-text.y)
			if err != nil {
				return err
			}

			if err := t.display.SetGlyph(ix, iy, gl.(Glyph)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Terminal) Render(surface Surface) error {
	return t.display.Render(surface, t.face)
}
