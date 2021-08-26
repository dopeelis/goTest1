package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем счетчик горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}

	// Вызываем функцию вывода через итерацию по элементам массива
	for _, i := range arr {
		wg.Add(1)
		go print(i, wg)
	}
	wg.Wait()
}

// Функция выводит значение в консоль
func print(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	wg.Done() // Убавляем счетчик на 1 после вывода

}
