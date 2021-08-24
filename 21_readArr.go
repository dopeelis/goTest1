package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	wg.Add(len(arr))
	go func() {
		for _, i := range arr {
			fmt.Println(i)
			wg.Done()
		}
	}()
	wg.Wait()
}
