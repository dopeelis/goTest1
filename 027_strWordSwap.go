package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println("Str 1:", str)
	fmt.Println("Str 2:", wordReverse(str))
}

func wordReverse(str string) (res string) {
	words := strings.Fields(str)
	for _, i := range words {
		res = string(i) + " " + res
	}
	return
}
