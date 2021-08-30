package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]string)

	mu := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	keyArr := []int{2, 4, 6, 8}
	valArr := []string{"two", "four", "six", "восемь"}

	for _, k := range keyArr {
		for _, v := range valArr {
			wg.Add(1)
			go func(m map[int]string, k int, v string, mu *sync.Mutex, wg *sync.WaitGroup) {
				mu.Lock()
				m[k] = v
				mu.Unlock()
				wg.Done()
			}(myMap, k, v, mu, wg)
			break
		}
		valArr = valArr[1:]
	}
	wg.Wait()

	mu.Lock()
	fmt.Println(myMap)
	mu.Unlock()

}
