package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Первый вариант
	b := bytes.Buffer{}

	// Второй вариант
	b2 := strings.Builder{}

	s := []string{"i", "am", "not", "a", "boy"}

	for _, i := range s {
		b.WriteString(i + " ")
	}

	fmt.Println("1:", b.String())

	for _, i := range s {
		b2.WriteString(i + " ")
	}

	fmt.Println("2:", b2.String())

	// Третий вариант
	fmt.Println("3:", strings.Join(s, " "))
}
