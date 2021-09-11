package primitives

import "fmt"

type (
	Array2D struct {
		bounds   Rect
		elements []interface{}
	}

	Array2DInitialiser = func() interface{}
)

func NewArray2D(width, height int, init Array2DInitialiser) Array2D {
	a := Array2D{Rect{Vector{0, 0}, Vector{width, height}}, make([]interface{}, width*height)}

	for i := 0; i < len(a.elements); i++ {
		a.elements[i] = init()
	}

	return a
}

func (a Array2D) Get(x, y int) (interface{}, error) {
	if err := a.checkBounds(x, y); err != nil {
		return nil, err
	}

	return a.elements[y*a.Width()+x], nil
}

func (a *Array2D) Set(x, y int, value interface{}) error {
	if err := a.checkBounds(x, y); err != nil {
		return err
	}

	a.elements[y*a.Width()+x] = value

	return nil
}

func (a Array2D) Width() int {
	return a.bounds.Width()
}

func (a Array2D) Height() int {
	return a.bounds.Height()
}

func (a Array2D) Size() Vector {
	return a.bounds.size
}

func (a Array2D) checkBounds(x, y int) error {
	if (x < 0 || x >= a.Width()) || (y < 0 || y >= a.Height()) {
		return a.outOfBounds(x, y, a.Size())
	}

	return nil
}

func (a Array2D) outOfBounds(x, y int, size Vector) error {
	return fmt.Errorf("x: %d, y: %d -> out of bounds, (size x: %d, y: %d)", x, y, size.x, size.y)
}
