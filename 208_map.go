package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 124
	m[2] = 281

	fmt.Println(m)

	// Не гарантируется одинаковый вывод при каждом вызове!
	for _, i := range m {
		fmt.Println(i)
	}
}
