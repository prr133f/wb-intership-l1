package obj

import "math"

type Point struct {
	x float32
	y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float32 {
	return float32(math.Sqrt(math.Pow(float64(other.x-p.x), 2) + math.Pow(float64(other.y-p.y), 2)))
}
