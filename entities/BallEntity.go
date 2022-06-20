package entities

import (
	"image/color"

	"github.com/airtonix/rpg/components"
)

// Your game object
type Ball struct {
	components.Control  // Ball is controlled by player
	components.Position // Ball position
	components.Velocity // Ball velocity
	components.Size     // Ball radius
	components.Apperance
}

func NewBall(x int, y int, radius int, speed float64, color color.Color) *Ball {
	ball := &Ball{
		components.NewControl(1),
		components.NewPosition(x, y),
		components.NewVelocity(0, 0),
		components.Size{
			W: float64(radius),
			H: float64(radius),
		},
		components.NewAppearance(color),
	}
	return ball
}
