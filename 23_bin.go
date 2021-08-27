package main

import "fmt"

func main() {
	// Задаем исходный массив (элементы д.б. отсортированы)
	arr := []int{2, 4, 5, 7, 10, 24, 56, 100, 120, 259, 300}
	// Число, которое ищем
	myNum := 4

	res := binarySearch(arr, myNum)

	// Вывод результата в консоль
	if res == -1 {
		fmt.Println("This number is not in the array")
	} else {
		fmt.Printf("My number %v, index: %v\n", myNum, res)
	}

}

func binarySearch(arr []int, obj int) int {
	// Задаем значение начального индекса
	low := 0

	// Задаем значение конечного индекса
	high := len(arr) - 1

	// Задаем значение индекса середины
	mid := len(arr) / 2

	for low <= high {

		// Проверяем середину
		if arr[mid] == obj {
			return mid
		}

		// Если во "второй" половине нет, то меняем значение вернего индекса
		// Обновляем значение индекса середины
		if arr[mid] > obj {
			high = mid - 1
			mid = (low + high) / 2
			continue
		}

		// Если в "первой" половине нет, то меняем значение нижнего индекса
		low = mid + 1
		// Обновляем значение индекса середины
		mid = (low + high) / 2
	}

	return -1
}
