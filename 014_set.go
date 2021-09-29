package main

import "fmt"

func main() {
	// Задаем исходные набор строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(arr)

	// Имитируем setM через map
	setM := make(map[string]struct{})

	// Задаем значение struct{}
	var exists = struct{}{}

	// Заполняем мап данными
	for _, i := range arr {
		setM[i] = exists
	}
	fmt.Println(setM)

	// Создаем слайс
	set := []string{}

	// Записываем значения из нашей мапы в слайс
	for i := range setM {
		set = append(set, i)
	}

	fmt.Println(set)
}
