package graphics

type (
	Terminal struct {
		width, height int
		writer        GlyphWriter
		bg, fg        Color
	}

	GlyphWriter interface {
		WriteGlyph(x, y int, g Glyph)
	}
)

func NewTerminal(width, height int, writer GlyphWriter) Terminal {
	return Terminal{width, height, writer, Black, White}
}

func (t Terminal) Width() int {
	return t.width
}

func (t Terminal) Height() int {
	return t.height
}

func (t *Terminal) Fill(x, y, width, height int, color Color) {
	glyph := NewGlyphFromCharCode(Space, t.fg, color)

	for iy := y; iy < y+height; iy++ {
		for ix := x; ix < x+width; ix++ {
			t.writer.WriteGlyph(ix, iy, glyph)
		}
	}
}

func (t *Terminal) WriteChar(x, y int, c CharCode, fg, bg Color) {
	t.writer.WriteGlyph(x, y, Glyph{c, fg, bg})
}

func (t *Terminal) WriteString(x, y int, text string, fg, bg Color) {
	runes := []rune(text)

	for i := range runes {
		if x+i >= t.width {
			break
		}

		t.writer.WriteGlyph(x+i, y, NewGlyphFromRune(runes[i], fg, bg))
	}
}
