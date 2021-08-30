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

	// Создаем буферизированный канал
	// Пустая структура занимает меньше памяти
	ch := make(chan struct{}, end)

	// Создем счетчики горутин
	wg := sync.WaitGroup{}
	wg2 := sync.WaitGroup{}

	wg.Add(end)
	// Наполняем канал пустыми структурами
	for i := 0; i < end; i++ {
		go func(i int) {
			ch <- struct{}{}
			wg.Done()
		}(i)
	}

	wg2.Add(1)
	// Увеличиваем значение счетчика
	go func() {
		for range ch {
			counter++
		}
		wg2.Done()
	}()

	wg.Wait()
	close(ch)
	wg2.Wait()

	fmt.Println(counter)
}
