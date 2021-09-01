package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем счетчик для горутин
	wg := new(sync.WaitGroup)

	// Создаем мьютекс
	mu := new(sync.Mutex)

	// Задаем исходные данные
	arr := [5]int{2, 4, 6, 8, 10}

	// Объясвляем переменную, куда будет записан результат
	res := 0
	// И ссылку на нее для работы в функциях
	resP := &res

	// Проходимся циклом по массиву
	for _, i := range arr {
		wg.Add(1)
		go func(res *int, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			// Считаем квадрат поступившего числа
			sq := i * i
			// Блокируем доступ
			mu.Lock()
			// Прибавляем это значение к результату
			*res += sq
			// "Отпускаем"
			mu.Unlock()
		}(resP, i, mu, wg)

	}
	// Ждем окончания выполнения всех горутин
	wg.Wait()

	// Выводим результат в консоль
	fmt.Println(res)
}
