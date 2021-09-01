package main

import (
	"fmt"
	"sync"
)

// Выбираем количество воркеров
const workerNum = 5

func main() {
	// Создание счетчика для горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные (можно использовать рандом, но так нагляднее)
	arr := []int{3, 6, 8, 10, 12, 15, 49, 38, 1, 23, 3, 7}
	wg.Add(len(arr))

	// Создаем пустой канал, куда будем записывать числа
	randChan := make(chan int, 2)

	// Запускаем работу воркеров
	for w := 1; w < workerNum+1; w++ {
		go worker(w, randChan, wg)
	}

	// Записываем в канал числа из массива
	for _, i := range arr {
		randChan <- i
	}

	// Завершаем работу всех воркеров
	close(randChan)
	// Можно добавить число для остановки воркера
	// Их количетсво равно количеству воркеров

	// Ждем завершения всех горутин
	wg.Wait()

}

// Функция, выводящая в консоль значения из канала с помощью воркеров
func worker(id int, ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		defer wg.Done()
		fmt.Println("Worker", id, "started")
		fmt.Println(i)
	}
	fmt.Println("worker", id, "finished")
}
