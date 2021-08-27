package main

import "fmt"

func main() {
	// Задаем исходные набор строк
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(arr)

	// Имитируем set через map
	set := make(map[string]struct{})

	// Задаем значение struct{}
	var exists = struct{}{}

	// Заполняем мап данными
	for _, i := range arr {
		set[i] = exists
	}
	fmt.Println(set)
}
