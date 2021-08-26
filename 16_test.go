package main

import "fmt"

func main() {
	n := 0 // (1)
	if true {
		n := 1 // Создается локальная переменная (внутри блока if) (2)
		n++
		fmt.Println("in func", n) // Выводится значение локальной переменной (2)
	}
	fmt.Println("out of func", n) // Выводится переменная (1)
}
