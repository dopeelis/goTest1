package main

import (
	"firstMod/traning/31_point/point"
	"fmt"
)

func main() {
	// // Объявляем первую точку
	point1 := point.Init(3, 2)

	// // Объявляем первую точку
	point2 := point.Init(5, 6)

	// Считаем результат и ыводим в консоль
	res := point.PointDistance(point2, point1)
	fmt.Printf("Distance between point 1 and point 2: %f\n", res)
}
