package graphics_test

import (
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/stretchr/testify/assert"
	"testing"
)

type screen struct {
	output  [][]graphics.Glyph
}

func newScreen(width, height int) screen {
	buf := make([][]graphics.Glyph, width)

	for x := 0; x < width; x++ {
		buf[x] = make([]graphics.Glyph, height)
		for y := 0; y < height; y++ {
			buf[x][y] = graphics.Clear
		}
	}

	return screen{buf}
}

func (s *screen) WriteGlyph(x, y int, g graphics.Glyph) {
	s.output[x][y] = g
}

func TestNewTerminal(t *testing.T) {
	const width, height = 10, 10

	monitor := newScreen(width, height)
	terminal := graphics.NewTerminal(width, height, &monitor)

	assert.Equal(t, width, terminal.Width())
	assert.Equal(t, height, terminal.Height())

	// NOTE (c.nicola): First we fill the screen completely with a new background
	terminal.Fill(0, 0, width, height, graphics.Aqua)

	for x := range monitor.output {
		for y := range monitor.output[x] {
			assert.Equal(t, graphics.Aqua, monitor.output[x][y].BackgroundColor())
		}
	}

	// NOTE (c.nicola): We write a single char to our terminal
	terminal.WriteChar(5, 5, graphics.Minus, graphics.Gold, graphics.Brown)
	assert.Equal(t, graphics.NewGlyphFromCharCode(graphics.Minus, graphics.Gold, graphics.Brown), monitor.output[5][5])
	// NOTE (c.nicola): Now we write a string to the terminal. The string fits nicely and
	//					won't be cut off.
	const hi = "Hi!"

	terminal.WriteString(0, 0, hi, graphics.LightGreen, graphics.White)

	for x := 0; x < len(hi); x++ {
		assert.Equal(t, graphics.NewGlyphFromRune(rune(hi[x]), graphics.LightGreen, graphics.White), monitor.output[x][0])
	}
	// NOTE (c.nicola): Now we write a string to the terminal that does not fit. The string will be
	//					cut off, it won't wrap to next line.
	const largeStr = "This is a large string"

	terminal.WriteString(0, 0, largeStr, graphics.Red, graphics.Yellow)

	for x := 0; x < width; x++ {
		assert.Equal(t, graphics.NewGlyphFromRune(rune(largeStr[x]), graphics.Red, graphics.Yellow), monitor.output[x][0])
		// NOTE (c.nicola): We check the second row in our frame - we expect that nothing is
		//					written there, fill color should be equal to what we filled in our
		//					first render call
		assert.Equal(t, graphics.NewGlyphFromCharCode(graphics.Space, graphics.White, graphics.Aqua), monitor.output[x][1])
	}
}
