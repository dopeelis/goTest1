package main

import (
	"fmt"
)

var justString string

func main() {
	someFunc()
	fmt.Println(justString)
}

func someFunc() {
	v := createHugeString(1 << 10)
	// В памяти выделено место под оставшиеся значения, которые не используются
	justString = v[:100] // Переопределяем значение глобальной переменной
	// Это сложно отследить
}

func createHugeString(size int) string {
	var v string
	for i := 0; i < size; i++ {

		v += "A"
	}
	return v
}
