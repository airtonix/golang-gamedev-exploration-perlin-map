package systems

import (
	"fmt"
	"image/color"
	"math"

	"github.com/airtonix/rpg/components"
	"github.com/airtonix/rpg/core/array"
	"github.com/airtonix/rpg/core/perlinmap"
	"github.com/airtonix/rpg/core/position"
	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
)

/*
Renders chunks in the radius around a PerlinCameraPoint
*/
type PerlinMapSystem struct {
	world            *engine.World
	tilePixelSize    int // width and height of a tile
	chunkTileSize    int // width and height of a chunk
	viewRadiusChunks int // spread out this many chunks to render
	generator        *perlin.Perlin
	terrainMapper    func(value float64) color.RGBA
}

func NewPerlinMapSystem(
	world *engine.World,
	tilePixelSize int,
	chunkTileSize int,
	viewRadiusChunks int,
	terrainMapper func(value float64) color.RGBA,
) *PerlinMapSystem {
	perlinNoise := perlin.NewPerlin(2, 2, 3, 100)

	return &PerlinMapSystem{
		world:            world,
		tilePixelSize:    tilePixelSize,
		chunkTileSize:    chunkTileSize,
		viewRadiusChunks: viewRadiusChunks,
		generator:        perlinNoise,
		terrainMapper:    terrainMapper,
	}
}

func (system *PerlinMapSystem) Draw(world engine.World, screen *ebiten.Image) {
	// Get controlled objects
	entities := world.View(
		components.PerlinView{},
		components.Position{},
	)

	debugUi, _ := world.View(
		components.DebugText{},
	).Get()

	var debugText *components.DebugText
	debugUi.Get(&debugText)

	chunkSizePixel := system.tilePixelSize * system.chunkTileSize

	entities.Each(func(entity engine.Entity) {
		var view *components.PerlinView
		var pos *components.Position
		entity.Get(&view, &pos)

		/**
		 *
		 *  Spawn chunks
		 *
		 */
		chunkX := int(math.Floor(pos.Point.X / float64(chunkSizePixel)))
		chunkY := int(math.Floor(pos.Point.Y / float64(chunkSizePixel)))
		currentChunkId := hashVector(chunkX, chunkY)

		startX := chunkX - system.viewRadiusChunks
		startY := chunkY - system.viewRadiusChunks
		maxX := chunkX + system.viewRadiusChunks
		maxY := chunkY + system.viewRadiusChunks

		debugText.Content["chunk.sizePixel"] = fmt.Sprintf(
			"%d", chunkSizePixel)
		debugText.Content["chunk.startXY"] = fmt.Sprintf(
			"%d/%d", startX, startY)
		debugText.Content["chunk.startXYpixels"] = fmt.Sprintf(
			"%d/%d", startX*chunkSizePixel, startY*chunkSizePixel)
		debugText.Content["chunk.maxXY"] = fmt.Sprintf(
			"%d/%d", maxX, maxY)
		debugText.Content["chunk.maxXYpixels"] = fmt.Sprintf(
			"%d/%d", maxX*chunkSizePixel, maxY*chunkSizePixel)
		debugText.Content["chunk.id"] = currentChunkId
		debugText.Content["view.PosXY"] = fmt.Sprintf(
			"%2.f/%2.f", pos.Point.X, pos.Point.Y)

		for x := startX; x < maxX; x++ {
			for y := startY; y < maxY; y++ {

				posX := x * chunkSizePixel
				posY := y * chunkSizePixel
				chunkId := hashVector(x, y)
				_, hasChunk := view.Chunks[chunkId]

				if !hasChunk {
					chunk := perlinmap.NewChunk(chunkId, posX, posY)
					chunk.Generate(
						system.tilePixelSize,
						system.chunkTileSize,
						system.generator.Noise2D,
					)
					view.Chunks[chunkId] = chunk
				}
			}
		}

		/**
		 *
		 *  Render Chunks in range
		 *
		 */
		max := float64(chunkSizePixel * system.viewRadiusChunks)

		debugText.Content["view.ChunkCount"] = fmt.Sprintf("%d",
			(len(view.Chunks)))
		debugText.Content["view.radiusPixels"] = fmt.Sprintf("%.2f", max)
		debugText.Content["view.radius"] = fmt.Sprintf("%d",
			(system.viewRadiusChunks))

		chunksInRange := []perlinmap.Chunk{}
		var closestChunk perlinmap.Chunk
		var closestChunkDistance float64

		for _, chunk := range view.Chunks {
			distance := chunk.Distance(pos.Point)

			if closestChunkDistance == 0 {
				closestChunkDistance = distance
				closestChunk = chunk
			} else if distance < closestChunkDistance {
				closestChunkDistance = distance
				closestChunk = chunk
			}

			if distance <= max {
				chunksInRange = append(chunksInRange, chunk)

				alpha := uint8(128)
				if chunk.Id != currentChunkId {
					alpha = uint8(255)

				}
				// ebitenutil.DrawRect(
				// 	screen,
				// 	chunk.Position.X+2,
				// 	chunk.Position.Y+2,
				// 	float64(chunkSizePixel-4),
				// 	float64(chunkSizePixel-4),
				// 	color.RGBA{
				// 		R: colornames.Amber100.R,
				// 		G: colornames.Amber100.G,
				// 		B: colornames.Amber100.B,
				// 		A: alpha,
				// 	},
				// )
				array.VectorEach(
					chunk.Values,
					system.chunkTileSize,
					func(x int, y int, item float64) {
						col := system.terrainMapper(item)

						ebitenutil.DrawRect(
							screen,
							chunk.Position.X+float64(x*system.chunkTileSize),
							chunk.Position.Y+float64(y*system.chunkTileSize),
							float64(system.tilePixelSize),
							float64(system.tilePixelSize),
							color.RGBA{
								R: col.R,
								G: col.G,
								B: col.B,
								A: alpha,
							},
						)
					},
				)
			}
		}

		debugText.Content["chunksInRange"] = fmt.Sprintf("%d",
			len(chunksInRange))
		debugText.Content["clostestChunkDistance"] = fmt.Sprintf(
			"%2.f", closestChunkDistance)
		debugText.Content["closestChunk"] = fmt.Sprintf(
			"[%s] %2.f/%2.f",
			closestChunk.Id,
			closestChunk.Position.X,
			closestChunk.Position.Y)

	})

}

func hashVectorPosition(pos *position.Vector) string {
	return fmt.Sprintf("%f_%f", pos.X, pos.Y)
}
func hashVector(x int, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}
