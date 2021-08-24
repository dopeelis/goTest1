package main

import (
	"fmt"
)

func main() {
	// Способы объявление слайсов
	var slice1 = []int{0, 10, 20, 30, 40}

	var slice2 = make([]int, 5) // Длина == капасити
	// slice2 := make()
	var slice3 = make([]string, 5, 10) // Длина != капасити

	var slice4 = new([10]int)[0:5] // Длина = 10, капасити = 50
	// + срезами:
	slice5 := slice1[1:4]

	fmt.Println("1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("2:", slice2, "len:", len(slice2), "cap:", cap(slice2))
	fmt.Println("3:", slice3, "len:", len(slice3), "cap:", cap(slice3))
	fmt.Println("4:", slice4, "len:", len(slice4), "cap:", cap(slice4))
	fmt.Println("5:", slice5, "len:", len(slice5), "cap:", cap(slice5))

	// Способы объвления map
	var map1 map[string]int // Создание пустой мапы
	var map2 = make(map[string]int)
	// map2 := make()
	var map3 = make(map[string]int, 10) // Задаем емкость == 10
	map4 := map[string]int{"l": 1, "k": 2}

	fmt.Println("1:", map1, "len:", len(map1))
	fmt.Println("2:", map2, "len:", len(map2))
	fmt.Println("3:", map3, "len:", len(map3))
	fmt.Println("4:", map4, "len:", len(map4))

}
