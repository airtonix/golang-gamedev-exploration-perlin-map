package systems

import (
	"github.com/airtonix/rpg/components"
	"github.com/sedyh/mizu/pkg/engine"
)

// You can go through all entities that have a certain set of
// components specifying the requirements in the fields of the system
type MovementSystem struct {
	*components.Position // Current entity position
	*components.Velocity // Current entity velocity
}

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

// Apply velocity for each entity that has Pos and Vel
func (system *MovementSystem) Update(world engine.World) {
	// If they are registered components, they will not be nil
	system.Position.X += system.Velocity.X
	system.Position.Y += system.Velocity.Y
}
