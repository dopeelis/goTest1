package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем счетчик горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные
	arr := [5]int{2, 4, 6, 8, 10}

	// Вызываем функцию вывода через итерацию по элементам массива
	for _, i := range arr {
		wg.Add(1) //Прибавляем счетчик на 1
		go printSq(i, wg)
	}
	wg.Wait()
}

// Функция выводит значение в консоль
func printSq(i int, wg *sync.WaitGroup) int {
	defer wg.Done()
	res := i * i
	fmt.Println(res)
	// wg.Done() // Убавляем счетчик на 1 после вывода
	return res
}
