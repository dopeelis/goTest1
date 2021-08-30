package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем счетчик для горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные
	arr := [5]int{2, 4, 6, 8, 10}

	// Объясвляем переменную, куда будет записан результат
	res := 0
	// И ссылку на нее для работы в функциях
	resP := &res

	// Проходимся циклом по массиву
	for _, i := range arr {
		wg.Add(1)
		go func(res *int, i int, wg *sync.WaitGroup) {
			defer wg.Done()
			// Считаем квадрат поступившего числа
			sq := i * i
			// Прибавляем это значение к результату
			*res += sq
		}(resP, i, wg)

	}
	// Ждем окончания выполнения всех горутин
	wg.Wait()

	// Выводим результат в консоль
	fmt.Println(res)
}
