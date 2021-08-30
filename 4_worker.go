package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создание счетчика для горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные (можно использовать рандом, но так нагляднее)
	arr := []int{3, 6, 8, 10, 12, 15, 49, 38, 1, 23, 3, 7}
	// Выделяем для определения буфера канала и числа горутин
	l := len(arr)

	// Создаем пустой канал, куда будем записывать числа
	randChan := make(chan int, l)

	// Выбираем количество воркеров
	workerNum := 5

	wg.Add(l)

	// Записываем в канал числа из массива
	for _, i := range arr {
		randChan <- i
	}
	close(randChan)

	// Запускаем работу воркеров
	for w := 1; w < workerNum+1; w++ {
		go worker(w, randChan, wg)
	}

	// Ждем завершения всех горутин
	wg.Wait()

}

// Функция, выводящая в консоль значения из канала с помощью воркеров
func worker(id int, ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("Worker", id, "started")
		fmt.Println(i)
		fmt.Println("worker", id, "finished")
		wg.Done()
	}
}
