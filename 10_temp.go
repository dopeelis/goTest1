package main

import "fmt"

func main() {
	// Задаем значения температуры
	data := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(data) // Проверка

	// Создаем "категории" для температуры
	tempCat := []int{-30, -20, -10, 0, 10, 20, 30}

	// Создаем мап, в котором будут храниться распределнные температуры
	tempGroup := make(map[int][]float32)

	// Присваиваем ключам мапы "категории" температур
	// Присваиваем значениям мапы пустые слайсы
	for _, i := range tempCat {
		tempGroup[i] = []float32{}
	}
	fmt.Println(tempGroup) // Проверка

	// Цикл для распределения температур по "категориям"
	for k := range tempGroup {
		for _, i := range data {
			// Если подходит в границы нашей категории
			if i >= float32(k) && i <= float32(k+9) {
				// Добавляем ее в наш слайс значений
				tempGroup[k] = append(tempGroup[k], i)
			} else {
				// Если нет, то идем дальше
				continue
			}
		}
	}
	fmt.Println(tempGroup)
}
