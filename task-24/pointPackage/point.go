package pointPackage

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (pointA *Point) Distance(pointB *Point) float64 {
	distX := pointA.x - pointB.x
	distY := pointA.y - pointB.y
	return math.Sqrt(distX*distX + distY*distY)
}
