package distance

import (
	"errors"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) (point Point, err error) {
	if x < 0 || y < 0 {
		err = errors.New("must be greater than 0")
		return point, err
	}
	point.x = x
	point.y = y
	return point, err
}

func Calculate(p1, p2 Point) float64 {
	distance := math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
	return distance
}
