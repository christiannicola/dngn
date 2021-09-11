package graphics

var (
	Clear = Glyph{Space, White, Black}
)

type Glyph struct {
	character CharCode
	fg, bg    Color
}

func NewGlyphFromCharCode(c CharCode, fg Color, bg Color) Glyph {
	return Glyph{c, fg, bg}
}

func NewGlyphFromRune(r rune, fg Color, bg Color) Glyph {
	return Glyph{CharCode(r), fg, bg}
}

func (g Glyph) Equal(other Glyph) bool {
	return g.character == other.character && g.fg.Equal(other.fg) && g.bg.Equal(other.bg)
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

func (g Glyph) ForeGroundColor() Color {
	return g.fg
}

func (g Glyph) BackgroundColor() Color {
	return g.bg
}
