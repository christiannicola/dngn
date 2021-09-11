package primitives

import "math"

type Vector struct {
	x int
	y int
}

func (v Vector) X() int {
	return v.x
}

func (v Vector) Y() int {
	return v.y
}

func (v Vector) LengthSquared() int {
	return v.x*v.x + v.y*v.y
}

func (v Vector) Area() int {
	return v.x * v.y
}

func (v Vector) Length() float64 {
	return math.Sqrt(float64(v.LengthSquared()))
}

func (v Vector) KingLength() int {
	return int(math.Max(math.Abs(float64(v.x)), math.Abs(float64(v.y))))
}

func (v Vector) RookLength() int {
	return int(math.Abs(float64(v.x) * math.Abs(float64(v.y))))
}

func (v Vector) Abs() Vector {
	return Vector{int(math.Abs(float64(v.x))), int(math.Abs(float64(v.y)))}
}

func (v Vector) Offset(x, y int) Vector {
	return Vector{v.x + x, v.y + y}
}

func (v Vector) OffsetX(x int) Vector {
	return Vector{v.x + x, v.y}
}

func (v Vector) OffsetY(y int) Vector {
	return Vector{v.x, v.y + y}
}

func (v Vector) Add(rhs Vector) Vector {
	return Vector{v.x + rhs.x, v.y + rhs.y}
}

func (v Vector) Subtract(rhs Vector) Vector {
	return Vector{v.x - rhs.x, v.y - rhs.y}
}

func (v Vector) Greater(rhs Vector) bool {
	return v.LengthSquared() > rhs.LengthSquared()
}

func (v Vector) GreaterEqual(rhs Vector) bool {
	return v.LengthSquared() >= rhs.LengthSquared()
}

func (v Vector) Less(rhs Vector) bool {
	return v.LengthSquared() < rhs.LengthSquared()
}

func (v Vector) LessEqual(rhs Vector) bool {
	return v.LengthSquared() <= rhs.LengthSquared()
}

func (v Vector) Equal(rhs Vector) bool {
	return (v.x == rhs.x) && (v.y == rhs.y)
}

func (v Vector) RotateLeft45() Direction {
	switch v {
	case North:
		return NorthWest
	case NorthEast:
		return North
	case East:
		return NorthEast
	case SouthEast:
		return East
	case South:
		return SouthEast
	case SouthWest:
		return South
	case West:
		return SouthWest
	case NorthWest:
		return West
	case None:
		return None
	default:
		return None
	}
}

func (v Vector) RotateRight45() Direction {
	switch v {
	case North:
		return NorthEast
	case NorthEast:
		return East
	case East:
		return SouthEast
	case SouthEast:
		return South
	case South:
		return SouthWest
	case SouthWest:
		return West
	case West:
		return NorthWest
	case NorthWest:
		return North
	case None:
		return None
	default:
		return None
	}
}

func (v Vector) RotateLeft90() Direction {
	switch v {
	case North:
		return West
	case NorthEast:
		return NorthWest
	case East:
		return North
	case SouthEast:
		return NorthEast
	case South:
		return East
	case SouthWest:
		return SouthEast
	case West:
		return South
	case NorthWest:
		return SouthWest
	case None:
		return None
	default:
		return None
	}
}

func (v Vector) RotateRight90() Direction {
	switch v {
	case North:
		return East
	case NorthEast:
		return SouthEast
	case East:
		return South
	case SouthEast:
		return SouthWest
	case South:
		return West
	case SouthWest:
		return NorthWest
	case West:
		return North
	case NorthWest:
		return NorthEast
	case None:
		return None
	default:
		return None
	}
}

func (v Vector) Rotate180() Direction {
	switch v {
	case North:
		return South
	case NorthEast:
		return SouthWest
	case East:
		return West
	case SouthEast:
		return NorthWest
	case South:
		return North
	case SouthWest:
		return NorthEast
	case West:
		return East
	case NorthWest:
		return SouthEast
	case None:
		return None
	default:
		return None
	}
}
