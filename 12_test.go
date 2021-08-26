package main

import "fmt"

// Функция с ошибкой

// func update(p *int) {
// 	  b := 2
// 	  p = &b
// }

// func main() {
// 	var (
// 	   a = 1
// 	   p = &a
// 	)
// 	fmt.Println(*p) // Вернет 1
// 	update(p)
// 	fmt.Println(*p) // Тоже вернет 1
// }
//

// Рабочая функция

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println("first call", *p)  // Выводит 1
	p = update(p)                  // Переопределяем значение нашего указателя
	fmt.Println("second call", *p) // Проверка: выводит 5
}

func update(p *int) *int { // Должен возвращаться указатель
	fmt.Println("first func call", *p)
	b := 5
	p = &b
	fmt.Println("second func call", *p)
	return p
}
