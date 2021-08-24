package main

import "fmt"

func main() {
	str := "❶ ❷ ❸ ❹ ❺ ❻"
	fmt.Println("Str 1:", str)
	fmt.Println("Str 2:", reverse(str))
}

func reverse(str string) (res string) {
	for _, i := range str {
		res = string(i) + res
	}
	return
}
