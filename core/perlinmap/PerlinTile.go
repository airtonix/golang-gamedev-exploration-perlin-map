package perlinmap

type Tile struct {
	x     int
	y     int
	value float64
}

func NewTile(x int, y int, value float64) *Tile {
	return &Tile{x, y, value}
}
