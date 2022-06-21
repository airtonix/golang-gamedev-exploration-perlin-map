package game

import (
	"image/color"

	"github.com/airtonix/rpg/components"
	"github.com/airtonix/rpg/entities"
	"github.com/airtonix/rpg/systems"
	"github.com/sedyh/mizu/pkg/engine"
	"golang.org/x/exp/shiny/materialdesign/colornames"
)

type Game struct{}

// Main scene, you can use that for settings or main menu
func (game *Game) Setup(world engine.World) {
	world.AddComponents(
		components.PerlinView{},
		components.Control{},
		components.Apperance{},
		components.Position{},
		components.Velocity{},
		components.Size{},
		components.Text{},
		components.DebugText{},
	)

	world.AddSystems(
		systems.NewPerlinMapSystem(
			&world,
			8, 8, 3,
			func(perlinValue float64) color.RGBA {
				if perlinValue <= 0.4 {
					return colornames.Blue600
				} else if perlinValue <= 0.75 {
					return colornames.Green700
				}
				return colornames.White
			},
		),
		systems.NewKeyboardControlSystem(),
		systems.NewPlayerControlSystem(),
		systems.NewMovementSystem(),
		systems.NewRenderSystem(),
		systems.NewDebugTextRenderSystem(),
	)

	bounds := world.Bounds()

	world.AddEntities(
		entities.NewDebugHud(
			16, float64(bounds.Dy())-32,
			256, 64,
			colornames.Grey200,
			colornames.White,
		),
		entities.NewBall(
			bounds.Dx()/2,
			bounds.Dy()/2,
			10,
			1.8,
			colornames.Blue300,
		),
	)
}
