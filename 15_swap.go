package main

import "fmt"

func main() {
	x := 5
	y := 10
	fmt.Printf("x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x, y)

}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
