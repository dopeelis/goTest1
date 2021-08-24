package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Задаем исходные данные
	info := []string{"1", "re", "co", "ve", "ry"}
	// Создаем таймер для завершения программы
	_ = time.AfterFunc(time.Duration(10)*time.Second, func() {
		fmt.Println("Time is over")
		// Выходим из программы
		os.Exit(0)
	})

	// Читаем из канала и выводим в консоль
	ch := make(chan string)
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	// Записываем входные данные в канал
	go func() chan string {
		for _, i := range info {
			ch <- i
		}
		close(ch)
		return ch
	}()

}
