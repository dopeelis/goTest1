package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Создаем функцию, которая выдает рандомные значения
func randNum() int {
	return rand.Intn(100)
}

func main() {
	// Объявляем канал, куда будем записывать значения
	ch := make(chan int, 100)

	// Объявляем канал для завершения программы
	done := make(chan struct{})

	// Записываем значения в основной канал
	// Пока ждем записи из канала завершения
	go func() {
		for {
			select {
			case ch <- randNum():
			case <-done:
				close(ch)
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Функция для записи в канал завершения
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	// Выводим значения из основного канала
	for i := range ch {
		fmt.Println("value:", i)
	}

	// Обозначаем конец программы
	fmt.Println("finish")
}
