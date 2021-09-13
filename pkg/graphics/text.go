package graphics

import (
	"fmt"
	"github.com/christiannicola/dngn/pkg/primitives"
	"strconv"
	"strings"
)

const (
	Custom TextPosition = iota
	TopLeft
	TopCentered
	TopRight
	LeftCentered
	Centered
	RightCentered
	BottomLeft
	BottomCentered
	BottomRight
)

type (
	Text struct {
		glyphs primitives.Array2D
		x, y   int
		pos    TextPosition
		fg, bg Color
	}

	TextPosition int
)

func NewText(text string, fg, bg Color) *Text {
	lines := strings.Split(text, "\n")

	// NOTE (c.nicola): Calculating the size of the rectangle for the text
	height := len(lines)
	width := 0

	for i := range lines {
		if len(lines[i]) < width {
			continue
		}

		width = len(lines[i])
	}

	// NOTE (c.nicola): Pad the lines whose length is not equal the width of the rectangle
	for i := range lines {
		if len(lines[i]) != width {
			lines[i] = fmt.Sprintf("%-"+strconv.Itoa(width)+"v", lines[i])
		}
	}

	x, y := 0, 0
	init := func() interface{} {
		if x == width {
			x, y = 0, y+1
		}

		g := NewGlyphFromRune(rune(lines[y][x]), fg, bg)
		x++

		return g
	}

	return &Text{
		primitives.NewArray2D(width, height, init),
		0, 0, TopLeft, fg, bg,
	}
}

func (t *Text) SetX(x int) {
	t.pos = Custom
	t.x = x
}

func (t *Text) SetY(y int) {
	t.pos = Custom
	t.y = y
}

func (t Text) Size() primitives.Vector {
	return t.glyphs.Size()
}

func (t *Text) SetPos(pos TextPosition, x, y, rWidth, rHeight int) {
	// NOTE (c.nicola): We do nothing if custom position in passed in
	//					or if the text does not fit into the rectangle size
	if pos == Custom || rWidth < t.glyphs.Width() || rHeight < t.glyphs.Height() {
		return
	}

	rect := primitives.NewRect(x, y, rWidth, rHeight)

	var position primitives.Vector

	switch pos {
	case TopLeft:
		position = rect.TopLeft()
		t.x, t.y = position.X(), position.Y()
	case TopCentered:
		position = rect.Center()
		t.x, t.y = position.X()-(t.glyphs.Width()/2), rect.Top()
	case TopRight:
		position = rect.TopRight()
		t.x, t.y = position.X()-t.glyphs.Width(), position.Y()
	case LeftCentered:
		t.x, t.y = rect.X(), rect.Y()+(rect.Height()/2)+1-t.glyphs.Height()
	case Centered:
		position = rect.Center()
		t.x, t.y = position.X()-(t.glyphs.Width()/2), position.Y()-(t.glyphs.Height()/2)
	case RightCentered:
		t.x, t.y = rect.Right()-t.glyphs.Width(), rect.Y()+(rect.Height()/2)+1-t.glyphs.Height()
	case BottomLeft:
		position = rect.BottomLeft()
		t.x, t.y = position.X(), position.Y()-t.glyphs.Height()
	case BottomCentered:
		position = rect.Center()
		t.x, t.y = position.X()-(t.glyphs.Width()/2), rect.Bottom()-t.glyphs.Height()
	case BottomRight:
		position = rect.BottomRight()
		t.x, t.y = position.X()-t.glyphs.Width(), position.Y()-t.glyphs.Height()
	}
}
