package components

import (
	"github.com/airtonix/rpg/core/perlinmap"
)

/*
Represents the focus point of the game camera.

The generation of chunks should be within a radius around this point.
*/
type PerlinView struct {
	Chunks map[string]perlinmap.Chunk
}

func NewPerlinView() PerlinView {
	return PerlinView{
		Chunks: map[string]perlinmap.Chunk{},
	}
}
