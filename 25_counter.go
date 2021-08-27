package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем первое значение для счетчика
	counter := 0
	// Задаем конечное значение
	end := 200
	// Создаем "блокировщик" для потоков
	m := sync.Mutex{}
	// Создаем счетчик для потоков
	wg := sync.WaitGroup{}

	// Количство потоков равно конечному значению для счетчика
	wg.Add(end)
	for i := 0; i < end; i++ {
		go func(i int) {
			// Блокируем использование счетчика для других потоков
			m.Lock()
			counter++
			// "Отпускаем" счетчик
			m.Unlock()
			// Завершаем горутину
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println(counter)
}
