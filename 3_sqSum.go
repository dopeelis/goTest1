package main

import "fmt"

func main() {
	// Задаем исходные данные
	m := []int{2, 4, 6, 8, 10}
	// Выводим в консоль
	fmt.Println(sqSum(sq(sliceChan(m))))
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
	sqChan := make(chan int)
	go func() {
		for i := range ch {
			sqChan <- i * i
		}
		// Закрываем канал, когда все квадраты чисел были отправлены
		close(sqChan)
	}()
	return sqChan
}

// Читаем из канала с квадратами чисел
func sqSum(ch <-chan int) int {
	// Объявляем переменную, в которую будет записан результат
	res := 0
	// Объявляем путь до переменной, чтобы ссылаться внутри функции
	resP := &res
	for i := range ch {
		// Прибавляем к результату квадраты числел из канала
		*resP += i
	}
	return res
}
