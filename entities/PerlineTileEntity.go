package entities

import (
	"image/color"

	"github.com/airtonix/rpg/components"
)

type PerlinTileEntity struct {
	components.Position  // Tile position
	components.Size      // Tile size
	components.Apperance // Tile color
}

func NewPerlinTile(
	x float64,
	y float64,
	radius float64,
	color color.Color,
) *PerlinTileEntity {
	tile := &PerlinTileEntity{
		components.NewPositionF(x, y),
		components.NewSizeF(radius, radius),
		components.NewAppearance(color),
	}
	return tile
}
