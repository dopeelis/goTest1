package main

import "fmt"

func main() {
	// Объявляем оба массива
	arr1 := [7]string{"ab", "ba", "ca", "cda", "f", "f", "a"}
	arr2 := [5]string{"a", "fc", "ac", "a", "t"}

	// Выводим результат в консоль
	fmt.Println(arrCros(arr1, arr2))
}

// Функция, находящая пересечения между массивами
func arrCros(arr1 [7]string, arr2 [5]string) (res []string) {
	// Добавляем значения первого массива в мап
	mArr1 := make(map[string]bool)
	mArr2 := make(map[string]bool)

	// Избавляемся от повторений в массиве
	for _, i := range arr1 {
		mArr1[i] = true
	}

	// Избавляемся от повторений в массиве
	for _, i := range arr2 {
		mArr2[i] = true
	}

	// Проходимся циклом и, если находим значение, добавляем к результату
	for k := range mArr1 {
		if mArr2[k] {
			res = append(res, k)
		}
	}
	return res
}
