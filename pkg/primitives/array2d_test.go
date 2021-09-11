package primitives_test

import (
	"github.com/christiannicola/dngn/pkg/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArray2D(t *testing.T) {
	// NOTE (c.nicola): 200 rows and columns in a 2d plane
	const width, height = 200, 200

	iter := -1

	init := func() interface{} {
		iter++

		return iter
	}

	a := primitives.NewArray2D(width, height, init)

	assert.Equal(t, width, a.Width())
	assert.Equal(t, height, a.Height())

	v, err := a.Get(100, 100)

	assert.NoError(t, err)
	// NOTE (c.nicola): first element has the value 0, last element has the value x * a.Width() + y (=> 200 * 200)
	assert.Equal(t, 100*a.Width()+100, v.(int))

	// NOTE (c.nicola): get an element out of bounds
	_, err = a.Get(500, 100)
	assert.Error(t, err)
	_, err = a.Get(100, 500)
	assert.Error(t, err)

	// NOTE (c.nicola): set an element
	err = a.Set(100, 100, 0)
	assert.NoError(t, err)

	v, err = a.Get(100, 100)
	assert.NoError(t, err)
	assert.Equal(t, 0, v.(int))

	// NOTE (c.nicola): set an element out of bounds
	err = a.Set(500, 100, 0)
	assert.Error(t, err)

	// NOTE (c.nicola): access the last element, zero based index
	v, err = a.Get(width-1, height-1)
	assert.NoError(t, err)
	assert.Equal(t, width*height-1, v.(int))
}
