package graphics_test

import (
	"fmt"
	"github.com/christiannicola/dngn/pkg/graphics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDisplay(t *testing.T) {
	const width, height = 6, 6
	display := graphics.NewDisplay(width, height)

	var frameBuffer [width][height]graphics.Glyph

	drawFn := func(x, y int, glyph graphics.Glyph) error {
		frameBuffer[x][y] = glyph

		return nil
	}

	// NOTE (c.nicola): draw stuff to the display
	assert.NoError(t, display.SetGlyph(1, 1, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(1, 2, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(1, 3, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(2, 1, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(2, 2, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(2, 3, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(3, 1, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(3, 2, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.SetGlyph(3, 3, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))

	assert.NoError(t, display.Render(drawFn))

	for x := 1; x <= 3; x++ {
		for y := 1; y <= 3; y++ {
			assert.Equal(t, graphics.Asterisk, frameBuffer[x][y].CharCode())
		}
	}

	// NOTE (c.nicola): Second render pass: a glyph is drawn to the display, but the new glyph is equal to the old one
	//					-> the display does not change the frame buffer for unnecessary draw calls
	//					-> record the pointer address as string and compare them later after the draw call
	curPtrAddress := fmt.Sprintf("%p", &frameBuffer[1][3])
	assert.NoError(t, display.SetGlyph(1, 3, graphics.NewGlyphFromCharCode(graphics.Asterisk, graphics.White, graphics.Black)))
	assert.NoError(t, display.Render(drawFn))
	assert.Equal(t, curPtrAddress, fmt.Sprintf("%p", &frameBuffer[1][3]))
}
