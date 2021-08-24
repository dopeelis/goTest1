package main

import "fmt"

func main() {
	// Задаем исходные данные
	m := []int{2, 4, 6, 8, 10, 12}
	// Выводим в консоль, пройдясь циклом
	for i := range sq(sliceChan(m)) {
		fmt.Println(i)
	}
}

// Превращаем слайс исходных данных в канал
func sliceChan(arr []int) <-chan int {
	firstChan := make(chan int)
	go func() {
		for _, i := range arr {
			firstChan <- i
		}
		// Закрываем канал, когда все числа были отправлены
		close(firstChan)
	}()
	return firstChan
}

// Читаем числа из канала с исходными данными
// Записываем в новый канал квадрат каждого прочитанного числа
func sq(ch <-chan int) <-chan int {
	outChan := make(chan int)
	go func() {
		for i := range ch {
			outChan <- i * i
		}
		// Закрываем канал, когда все квадраты чисел были отправлены
		close(outChan)
	}()
	return outChan
}
