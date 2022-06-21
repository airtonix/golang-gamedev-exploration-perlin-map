package entities

import (
	"image/color"

	"github.com/airtonix/rpg/components"
)

type DebugHud struct {
	components.Position // Tile position
	components.Size     // Tile size
	components.DebugText
}

func NewDebugHud(
	x float64,
	y float64,
	w float64,
	h float64,
	bg color.Color,
	fg color.Color,
) *DebugHud {
	tile := &DebugHud{
		components.NewPositionF(x, y),
		components.NewSizeF(w, h),
		components.NewDebugText(fg),
	}
	return tile
}
