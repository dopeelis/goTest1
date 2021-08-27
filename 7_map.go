package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	// Слайс, из которого будем писать в мап
	lst := []int{2, 4, 6, 8}

	// Проходим циклом и в нем добавляем значения в мапу
	for i, v := range lst {
		// Объявлем счетчику о горутине
		wg.Add(1)
		go func(m map[int]int, k int, v int, mu *sync.Mutex, wg *sync.WaitGroup) {
			// Блокируем доступ других горутин к мапе
			mu.Lock()
			m[k] = v
			// Снова открываем доступ
			mu.Unlock()
			// Завершаем горутину
			wg.Done()
		}(myMap, i, v, mu, wg)
	}

	// Ждем выполнения всех горутин
	wg.Wait()

	mu.Lock()
	fmt.Println(myMap)
	mu.Unlock()

}
