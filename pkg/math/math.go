package math

func Clamp(x, min, max int) int {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}

func ClampUint8(x, min, max uint8) uint8 {
	return uint8(Clamp(int(x), int(min), int(max)))
}

func ClampFloat32(x, min, max float32) float32 {
	switch {
	case x < min:
		return min
	case x > max:
		return max
	default:
		return x
	}
}
