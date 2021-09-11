package graphics

type (
	Terminal struct {
		display Display
		bg, fg  Color
	}
)

func NewTerminal(width, height int) Terminal {
	return Terminal{NewDisplay(width, height), Black, White}
}

func (t Terminal) Width() int {
	return t.display.Width()
}

func (t Terminal) Height() int {
	return t.display.Height()
}

func (t *Terminal) Fill(x, y, width, height int, color Color) error {
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

func (t *Terminal) WriteChar(x, y int, c CharCode, fg, bg Color) error {
	return t.display.SetGlyph(x, y, Glyph{c, fg, bg})
}

func (t *Terminal) WriteString(x, y int, text string, fg, bg Color) error {
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

func (t *Terminal) Render(fn DrawGlyphFn) error {
	return t.display.Render(fn)
}