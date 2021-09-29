package main

import "fmt"

func main() {
	// Создаем новый слайс и задаем ему размер 100 элементов
	newSlice := make([]int, 100)
	// Проверяем capacity
	fmt.Println(cap(newSlice))
}
