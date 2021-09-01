package main

import (
	"fmt"
)

func main() {
	// Задаем значения температуры
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем мап, в котором будут храниться распределнные температуры
	res := make(map[int][]float32)

	for _, i := range data {
		// Добавляем "категории" температур
		category := int(i) - (int(i) % 10)
		// Если категория есть, то добавляем к ней значение температуры
		if _, ok := res[category]; ok {
			res[category] = append(res[category], i)
		} else {
			// Если нет,то создаем новую категорию и добавляем туда значение
			res[category] = []float32{i}
		}
	}
	fmt.Println(res)
}
