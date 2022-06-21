package position

import "math"

type Vector struct {
	X, Y float64
}

func NewVector(x, y int) *Vector {
	return &Vector{float64(x), float64(y)}
}
func NewVectorF(x, y float64) *Vector {
	return &Vector{x, y}
}

func (vector *Vector) Point() (float64, float64) {
	return vector.X, vector.Y
}

func (vector *Vector) Distance(to *Vector) float64 {
	first := math.Pow(vector.X-to.X, 2)
	second := math.Pow(vector.Y-to.Y, 2)
	return math.Sqrt(first + second)
}
