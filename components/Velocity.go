package components

// Velocity for any entity, if it needs
type Velocity struct {
	X, Y float64 // Also, 2D point
}

func NewVelocity(x, y int) Velocity {
	return Velocity{
		X: float64(x),
		Y: float64(y),
	}
}
