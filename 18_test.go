package main

import "fmt"

// Функция с ошибкой

// func someAction(v []int8, b int8) {
// 	v[0] = 100
// 	v = append(v, b)
//   }

//   func main() {
// 	var a = []int8{1, 2, 3, 4, 5}
// 	someAction(a, 6)
// 	fmt.Println(a)
//   }
//

// Работающая функция

func someAction(v []int8, b int8) []int8 { // Должен возвращаться слайс (если добавляем эл-т)
	v[0] = 100 // Меняется значение в оригинальном слайсе (изм. 1)
	fmt.Println("1", v, cap(v))
	v = append(v, b) // Создается новый массив и указывается на него (изм. 2)
	fmt.Println("2", v, cap(v))
	return v
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	a = someAction(a, 6) // Без присвоения будет только изменение 1, но не 2
	fmt.Println("3", a, cap(a))
}
