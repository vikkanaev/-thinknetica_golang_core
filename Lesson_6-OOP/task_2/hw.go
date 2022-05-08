package hw

import (
	"errors"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	x1, y1, x2, y2 float64
}

func (geom Geom) CalculateDistance() (distance float64, err error) {
	if geom.x1 < 0 || geom.x2 < 0 || geom.y1 < 0 || geom.y2 < 0 {
		err = errors.New("must be greater than 0")
		return distance, err
	}

	distance = math.Sqrt(math.Pow(geom.x2-geom.x1, 2) + math.Pow(geom.y2-geom.y1, 2))

	// возврат расстояния между точками и ошибку
	return distance, err
}
