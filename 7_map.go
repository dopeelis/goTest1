package main

import (
	"fmt"
	"sync"
)

func main() {
	// Объявляем мапу, в которую будем записывать значения
	myMap := make(map[int]string)

	// Объявляем локер
	mu := new(sync.Mutex)

	// Объявляем счетчик горутин
	wg := new(sync.WaitGroup)

	// Задаем исходные данные для ключей и значений
	keyArr := []int{2, 4, 6, 8}
	valArr := []string{"two", "four", "six", "восемь"}

	// Проходимся циклом по количеству значений и записываем их в мап
	for i := 0; i < len(keyArr); i++ {
		wg.Add(1)
		go func(m map[int]string, k []int, v []string, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			// Заблокировали доступ для других горутин
			mu.Lock()
			// Записали значение
			m[k[i]] = v[i]
			// "Отпустили"
			mu.Unlock()
		}(myMap, keyArr, valArr, i, mu, wg)
	}
	// Ждем выполнения всех горутин
	wg.Wait()

	// -race указывает, что тут возможно состояние гонки
	// Локи добавлены для перестраховки
	mu.Lock()
	fmt.Println(myMap)
	mu.Unlock()

}
