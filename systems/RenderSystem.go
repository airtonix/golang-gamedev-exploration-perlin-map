package systems

import (
	"github.com/airtonix/rpg/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

// When you need many sets of components
// in one system, you can use the views
type RenderSystem struct{}

func NewRenderSystem() *RenderSystem {
	return &RenderSystem{}
}

// Render one frame
func (system *RenderSystem) Draw(world engine.World, screen *ebiten.Image) {
	// But choose the right entities yourself
	view := world.View(
		components.Position{},
		components.Size{},
		components.Apperance{},
	)

	view.Each(func(entity engine.Entity) {
		var pos *components.Position
		var size *components.Size
		var appearance *components.Apperance

		entity.Get(&pos, &size, &appearance)

		ebitenutil.DrawRect(
			screen,
			pos.X-(size.W/2), pos.Y-(size.H/2),
			size.W, size.H,
			appearance.Color,
		)

	})
}
