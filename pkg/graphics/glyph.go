package graphics

import "image/color"

var (
	Clear = Glyph{Space, White, Black}
)

type Glyph struct {
	character CharCode
	fg, bg    color.Color
}

func NewGlyphFromCharCode(c CharCode, fg color.Color, bg color.Color) Glyph {
	return Glyph{c, fg, bg}
}

func NewGlyphFromRune(r rune, fg color.Color, bg color.Color) Glyph {
	return Glyph{CharCode(r), fg, bg}
}

func (g Glyph) Equal(other Glyph) bool {
	return g.character == other.character && AreColorsEqual(g.fg, other.fg) && AreColorsEqual(g.bg, other.bg)
}

func (g Glyph) Rune() rune {
	return rune(g.character)
}

func (g Glyph) CharCode() CharCode {
	return g.character
}

func (g Glyph) String() string {
	return string(g.Rune())
}

func (g Glyph) ForeGroundColor() color.Color {
	return g.fg
}

func (g Glyph) BackgroundColor() color.Color {
	return g.bg
}
