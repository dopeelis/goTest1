package main

import "fmt"

func main() {
	// Исходные данные
	x := 5
	y := 10
	fmt.Printf("x=%d, y=%d\n", x, y) // Вывод до изменения
	swap(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x, y) // Вывод после изменения

}

// Функция обмена значениями
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
