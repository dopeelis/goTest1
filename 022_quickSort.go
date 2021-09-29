package main

import "fmt"

func main() {
	arr := []int{7, 5, 24, 10, 4, 300, 259, 100, 120, 56}
	arr = quickSsort(arr)
	fmt.Println(arr)

}

func quickSsort(arr []int) []int {
	// Базовый случай
	if len(arr) < 2 {
		return arr
	}

	left := 0
	right := len(arr) - 1

	// Рекурсивый случай

	// Выбираем опорный элемент
	pivot := len(arr) / 2

	// Смещаем опорный элемент вправо
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Проходимся циклом по всем элементам
	for i := range arr {
		if arr[i] < arr[right] {
			// Если они меньше, смешаем их влево
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Меняем ячейки друг с другом
	// т.к. сам элемент arr[right] не был отсортирован
	// Теперь середина это left
	arr[left], arr[right] = arr[right], arr[left]

	// Проходимся сортировкой по оставшимся частям
	quickSsort(arr[:left])
	quickSsort(arr[left+1:])

	return arr
}
