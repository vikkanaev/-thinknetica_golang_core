package distance

import (
	"math"
)

// CalculateDistance возвращает растояние между двумя заданными точками
func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
