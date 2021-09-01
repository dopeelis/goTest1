package point

import (
	"math"
)

// Описываем характеристики точки
type point struct {
	x int
	y int
}

// Функция для определения значений точек
func Init(x int, y int) point {
	p := &point{}
	p.x = x
	p.y = y
	return *p
}

// Рассчитывем расстояние между точками по разнице координат
func PointDistance(p1 point, p2 point) float64 {
	delX := math.Abs(float64(p1.x - p2.x))
	delY := math.Abs(float64(p1.y - p2.y))
	distance := math.Sqrt(delX*delX + delY*delY)
	return distance
}
