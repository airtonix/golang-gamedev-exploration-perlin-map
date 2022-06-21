package systems

import (
	"fmt"
	"sort"

	"github.com/airtonix/rpg/components"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

// When you need many sets of components
// in one system, you can use the views
type DebugTextRenderSystem struct {
	*components.DebugText
}

func NewDebugTextRenderSystem() *DebugTextRenderSystem {
	return &DebugTextRenderSystem{}
}

// Render one frame
func (system *DebugTextRenderSystem) Draw(
	world engine.World,
	screen *ebiten.Image,
) {
	// But choose the right entities yourself
	view := world.View(
		components.DebugText{},
	)

	view.Each(func(entity engine.Entity) {
		var pos *components.Position
		var size *components.Size
		var appearance *components.Apperance
		var text *components.DebugText

		entity.Get(&pos, &size, &appearance, &text)

		height := int(size.H)
		keys := make([]string, 0, len(text.Content))
		for k := range text.Content {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			height += 16
			text := fmt.Sprintf("%s: %s", key, text.Content[key])

			ebitenutil.DebugPrintAt(
				screen, text,
				16,
				screen.Bounds().Dy()-height,
			)
		}

	})
}
