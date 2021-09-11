package graphics

import (
	"github.com/christiannicola/dngn/pkg/primitives"
)

type (
	Display struct {
		glyphs, changedGlyphs primitives.Array2D
	}

	DrawGlyphFn = func(x, y int, g Glyph)
)

func NewDisplay(width, height int) Display {
	initializer := func() interface{} { return Clear }

	return Display{primitives.NewArray2D(width, height, initializer), primitives.NewArray2D(width, height, initializer)}
}

func (d Display) Width() int {
	return d.glyphs.Width()
}

func (d Display) Height() int {
	return d.glyphs.Height()
}

func (d Display) Size() primitives.Vector {
	return d.glyphs.Size()
}

func (d *Display) SetGlyph(x, y int, glyph Glyph) error {
	if (x < 0) || (x >= d.Width()) || (y < 0) || (y >= d.Height()) {
		return nil
	}

	v, err := d.glyphs.Get(x, y)
	if err != nil {
		return err
	}

	vt, ok := v.(Glyph)
	if !ok {
		return ErrDisplayInvalidGlyph
	}

	if !vt.Equal(glyph) {
		return d.changedGlyphs.Set(x, y, glyph)
	}

	return d.changedGlyphs.Set(x, y, nil)
}

func (d *Display) Render(fn DrawGlyphFn) error {
	for y := 0; y < d.Height(); y++ {
		for x := 0; x < d.Width(); x++ {
			changed, err := d.changedGlyphs.Get(x, y)
			if err != nil {
				return err
			}

			if changed == nil {
				continue
			}

			changedGlyph, ok := changed.(Glyph)
			if !ok {
				return ErrDisplayInvalidGlyph
			}

			fn(x, y, changedGlyph)

			if err = d.glyphs.Set(x, y, changedGlyph); err != nil {
				return err
			}

			if err = d.changedGlyphs.Set(x, y, nil); err != nil {
				return err
			}
		}
	}

	return nil
}
