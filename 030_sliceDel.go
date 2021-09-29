package main

import "fmt"

func main() {
	// Задаем исходный слайс
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(arr)

	fmt.Println(deleteI(arr, 2))
}

// Функция удаления i-го элемента из слайса
func deleteI(sl []int, i int) []int {
	// Меняем местами выбранный элемент с последним
	sl[i], sl[len(sl)-1] = sl[len(sl)-1], sl[i]
	// Возвращаем слайс без последнего значения
	return sl[:len(sl)-1]
}
