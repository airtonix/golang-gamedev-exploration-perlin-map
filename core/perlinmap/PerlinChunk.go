package perlinmap

import (
	"github.com/airtonix/rpg/core/position"
)

type ChunkValueGenerator func(x float64, y float64) float64

type Chunk struct {
	Id       string
	IsLoaded bool
	Values   []float64
	Position position.Vector
}

func NewChunk(id string, x, y int) Chunk {
	return Chunk{
		Id:       id,
		Position: *position.NewVector(x, y),
	}
}

func (chunk *Chunk) Distance(to *position.Vector) float64 {
	return chunk.Position.Distance(to)
}

func (chunk *Chunk) Generate(
	size, tileSize int,
	generator ChunkValueGenerator,
) {
	if chunk.IsLoaded {
		return
	}
	chunkSize := float64(size * tileSize)
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			tileX := float64(x * tileSize)
			tileY := float64(y * tileSize)

			chunkX, chunkY := chunk.Position.Point()
			absoluteX := chunkX*chunkSize + tileX
			absoluteY := chunkY*chunkSize + tileY

			chunk.Values = append(
				chunk.Values,
				generator(absoluteX/1000.0, absoluteY/1000.0),
			)
		}
	}
	chunk.IsLoaded = true
}
