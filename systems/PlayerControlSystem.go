package systems

import (
	"github.com/airtonix/rpg/components"
	"github.com/airtonix/rpg/core/num"
	"github.com/sedyh/mizu/pkg/engine"
)

type PlayerControlSystem struct {
}

func NewPlayerControlSystem() *PlayerControlSystem {
	return &PlayerControlSystem{}
}

func (system *PlayerControlSystem) Update(world engine.World) {

	// Get controlled objects
	player, ok := world.View(
		components.Position{},
		components.Velocity{},
		components.Control{},
	).Get()

	if !ok {
		return
	}

	var pos *components.Position
	var vel *components.Velocity
	var control *components.Control

	player.Get(&pos, &vel, &control)

	moveDirectionX := 0.0
	moveDirectionY := 0.0

	if control.Direction["Right"] {
		moveDirectionX = control.MoveSpeed
	}

	if control.Direction["Left"] {
		moveDirectionX = -control.MoveSpeed
	}

	if control.Direction["Up"] {
		moveDirectionY = -control.MoveSpeed
	}

	if control.Direction["Down"] {
		moveDirectionY = control.MoveSpeed
	}

	vel.X = num.Lerp(vel.X, control.MoveSpeed*moveDirectionX, control.MoveSpeed)
	vel.Y = num.Lerp(vel.Y, control.MoveSpeed*moveDirectionY, control.MoveSpeed)

}
