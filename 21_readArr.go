package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем счетчик горутин
	wg := new(sync.WaitGroup)

	// Создаем мьютекс
	mu := new(sync.Mutex)

	// Задаем исходные данные
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// Счетчик ставим равным длине слайса
	wg.Add(len(arr))

	// Выводим в консоль по количеству элементов в слайсе
	for i := 0; i < len(arr); i++ {
		go func(arr []int, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
			// Отложить завершение горутины
			defer wg.Done()
			// Блокируем доступ остальных
			mu.Lock()
			// Выводим
			fmt.Println(arr[i])
			// "Отпускаем"
			mu.Unlock()
		}(arr, i, mu, wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()
}
