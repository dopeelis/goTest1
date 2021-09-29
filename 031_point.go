package main

import (
	"fmt"
	"math"
)

// Описываем характеристики точки
type point struct {
	x int
	y int
}

// Функция для определения значений точек
func (p *point) Init(x int, y int) {
	p.x = x
	p.y = y
}

func main() {
	// Объявляем первую точку
	point1 := new(point)
	// Задаем значение первой точки
	point1.Init(3, 2)
	fmt.Printf("Point 1 {x: %d, y: %d}\n", point1.x, point1.y)
	// Объявляем вторую точку
	point2 := new(point)
	// Задаем значение второй точки
	point2.Init(5, 6)
	fmt.Printf("Point 2 {x: %d, y: %d}\n", point2.x, point2.y)
	// Считаем результат и ыводим в консоль
	res := pointDistance(*point2, *point1)
	fmt.Printf("Distance between point 1 and point 2: %f\n", res)
}

// Рассчитывем расстояние между точками по разнице координат
func pointDistance(p1 point, p2 point) float64 {
	delX := math.Abs(float64(p1.x - p2.x))
	delY := math.Abs(float64(p1.y - p2.y))
	distance := math.Sqrt(delX*delX + delY*delY)
	return distance
}
