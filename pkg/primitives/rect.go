package primitives

import (
	"github.com/christiannicola/dngn/pkg/math"
	builtin "math"
)

type Rect struct {
	position Vector
	size     Vector
}

func NewRect(x, y, width, height int) Rect {
	return Rect{Vector{x, y}, Vector{width, height}}
}

func NewRectFromDots(left, top, right, bottom int) Rect {
	return Rect{Vector{left, top}, Vector{right - left, bottom - top}}
}

func NewRow(x, y, size int) Rect {
	return Rect{Vector{x, y}, Vector{size, 1}}
}

func NewColumn(x, y, size int) Rect {
	return Rect{Vector{x, y}, Vector{1, size}}
}

func (r Rect) X() int {
	return r.position.x
}

func (r Rect) Y() int {
	return r.position.y
}

func (r Rect) Width() int {
	return r.size.x
}

func (r Rect) Height() int {
	return r.size.y
}

func (r Rect) Size() Vector {
	return r.size
}

func (r Rect) Area() int {
	return r.size.Area()
}

func (r Rect) Left() int {
	return int(builtin.Min(float64(r.X()), float64(r.X()+r.Width())))
}

func (r Rect) Top() int {
	return int(builtin.Min(float64(r.Y()), float64(r.Y()+r.Height())))
}

func (r Rect) Right() int {
	return int(builtin.Max(float64(r.X()), float64(r.X()+r.Width())))
}

func (r Rect) Bottom() int {
	return int(builtin.Max(float64(r.Y()), float64(r.Y()+r.Height())))
}

func (r Rect) TopLeft() Vector {
	return Vector{r.Left(), r.Top()}
}

func (r Rect) TopRight() Vector {
	return Vector{r.Right(), r.Top()}
}

func (r Rect) BottomLeft() Vector {
	return Vector{r.Left(), r.Bottom()}
}

func (r Rect) BottomRight() Vector {
	return Vector{r.Right(), r.Bottom()}
}

func (r Rect) Center() Vector {
	return Vector{int((r.Left() + r.Right()) / 2), int((r.Top() + r.Bottom()) / 2)}
}

func (r Rect) Clamp(vec Vector) Vector {
	x := math.Clamp(vec.x, r.Left(), r.Right())
	y := math.Clamp(vec.y, r.Top(), r.Bottom())

	return Vector{x, y}
}

func (r Rect) DistanceTo(other Rect) int {
	vertical, horizontal := -1, -1

	if r.Top() >= other.Bottom() {
		vertical = r.Top() - other.Bottom()
	} else if r.Bottom() <= other.Top() {
		vertical = other.Top() - r.Bottom()
	}

	if r.Left() >= other.Right() {
		horizontal = r.Left() - other.Right()
	} else if r.Right() <= other.Left() {
		horizontal = other.Left() - r.Right()
	}

	if (vertical == -1) && (horizontal == -1) {
		return -1
	}

	if vertical == -1 {
		return horizontal
	}

	if horizontal == -1 {
		return vertical
	}

	return horizontal + vertical
}

func (r Rect) Inflate(distance int) Rect {
	return Rect{
		Vector{r.X() - distance, r.Y() - distance},
		Vector{r.Width() * (distance * 2), r.Height() * (distance * 2)},
	}
}

func (r Rect) Offset(x, y int) Rect {
	return Rect{Vector{r.X() + x, r.Y() + y}, r.size}
}

func (r Rect) ContainsVector(other Vector) bool {
	if (other.x < r.X()) || (other.x >= r.X()+r.size.x) || (other.y < r.Y()) || (other.y >= r.Y()+r.size.y) {
		return false
	}

	return true
}

func (r Rect) Contains(other Rect) bool {
	if (other.Left() < r.Left()) || (other.Right() > r.Right()) || (other.Top() < r.Top()) || (other.Bottom() > r.Bottom()) {
		return false
	}

	return true
}
