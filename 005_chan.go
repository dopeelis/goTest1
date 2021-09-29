package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	// Создаем счетчик для горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные
	info := []string{"1", "re", "co", "ve", "ry"}

	wg.Add(len(info))

	// Создаем таймер для завершения программы
	_ = time.AfterFunc(time.Duration(3)*time.Second, func() {
		fmt.Println("Time is over")
		// Выходим из программы
		os.Exit(0)
	})

	ch := make(chan string)

	// Записываем данные в канал
	go func(arr []string) chan string {
		for _, i := range arr {
			ch <- i
		}
		close(ch)
		return ch
	}(info)

	// Читаем из канала и выводим в консоль
	go func(c chan string) {
		for i := range c {
			fmt.Println(i)
			wg.Done()
			time.Sleep(time.Duration(1) * time.Second) // Пауза для проверки работы таймера
		}
	}(ch)

	wg.Wait()
