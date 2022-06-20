package components

import (
	"image/color"
)

type Apperance struct {
	Color color.Color
}

func NewAppearance(color color.Color) Apperance {
	return Apperance{
		Color: color,
	}
}
