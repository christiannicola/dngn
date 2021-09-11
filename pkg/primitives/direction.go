package primitives

var _ Direction = Vector{}

var (
	None      = Vector{0, 0}
	North     = Vector{0, -1}
	NorthEast = Vector{1, -1}
	East      = Vector{1, 0}
	SouthEast = Vector{1, 1}
	South     = Vector{0, 1}
	SouthWest = Vector{-1, 1}
	West      = Vector{-1, 0}
	NorthWest = Vector{-1, -1}

	AllDirections           = []Vector{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
	CardinalDirections      = []Vector{North, East, South, West}
	InterCardinalDirections = []Vector{NorthEast, SouthEast, NorthWest, SouthWest}
)

type Direction interface {
	RotateLeft45() Direction
	RotateRight45() Direction
	RotateLeft90() Direction
	RotateRight90() Direction
	Rotate180() Direction
}
