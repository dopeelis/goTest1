package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a") // Сразу создается ссылка на новый массив
		slice[0] = "b"             // Изменения вносятся уже в новый слайс
		slice[1] = "b"
		fmt.Println(slice, cap(slice))
	}(slice)
	fmt.Println(slice, cap(slice)) // Старый слайс не изменяется
}
