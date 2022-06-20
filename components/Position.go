package components

// Position for any entity, if it needs
type Position struct {
	X, Y float64 // Just a 2D point
}

func NewPosition(x, y int) Position {
	return Position{float64(x), float64(y)}
}
