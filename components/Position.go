package components

import "github.com/airtonix/rpg/core/position"

// Position for any entity, if it needs
type Position struct {
	Point *position.Vector
}

func NewPositionI(x, y int) Position {
	return Position{
		Point: position.NewVector(x, y),
	}
}
func NewPositionF(x, y float64) Position {
	return Position{
		Point: position.NewVectorF(x, y),
	}
}
