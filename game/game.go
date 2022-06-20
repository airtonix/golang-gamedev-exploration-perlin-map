package game

import (
	"github.com/airtonix/rpg/components"
	"github.com/airtonix/rpg/entities"
	"github.com/airtonix/rpg/systems"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/image/colornames"
)

type Game struct{}

// Main scene, you can use that for settings or main menu
func (game *Game) Setup(world engine.World) {
	world.AddComponents(
		components.Control{},
		components.Apperance{},
		components.Position{},
		components.Velocity{},
		components.Size{},
	)

	bounds := world.Bounds()

	world.AddEntities(
		entities.NewBall(
			bounds.Dx()/2,
			bounds.Dy()/2,
			10,
			2.2,
			colornames.Royalblue,
		),
	)
	world.AddSystems(
		systems.NewKeyboardControlSystem(),
		systems.NewPlayerControlSystem(),
		systems.NewMovementSystem(),
		systems.NewRenderSystem(),
	)
}
