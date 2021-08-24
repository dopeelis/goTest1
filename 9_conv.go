package main

import (
	"fmt"
	"sync"
)

func main() {
	// Задаем исходные данные
	myArr := []int{2, 4, 5, 6, 8, 10}
	// Объявляем счетчик исполняемых горутин
	var wg sync.WaitGroup
	// Задаем количество вызовов, равное количеству элементов
	wg.Add(len(myArr))
	// Выводим результат в консоль
	for i := range doubleChan(writeChan(myArr)) {
		fmt.Println(i)
		wg.Done()
	}
	wg.Wait()
}

// Записываем значения из слайса в канал
func writeChan(arr []int) chan int {
	outChan := make(chan int)
	go func() {
		for _, i := range arr {
			outChan <- i
		}
		close(outChan)
	}()
	return outChan
}

// Читаем значения из канала и записываем в новый х2
func doubleChan(ch chan int) chan int {
	outChan := make(chan int)
	go func() {
		for i := range ch {
			outChan <- 2 * i
		}
		close(outChan)
	}()
	return outChan
}
