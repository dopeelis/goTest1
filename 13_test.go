package main

import (
	"fmt"
	"sync"
)

// Функция с ошибкой

// func main() {
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 5; i++ {
// 	   wg.Add(1)
// 	   go func(wg sync.WaitGroup, i int) {
// 		  fmt.Println(i)
// 		  wg.Done()
// 	   }(wg, i)
// 	}
// 	wg.Wait()
// 	fmt.Println("exit")
//  }
//

// Рабочая функция

func main() {
	wg := sync.WaitGroup{} // (*) или wg :=new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) { // Нужно передавать ссылку на счетчик
			fmt.Println(i)
			wg.Done()
		}(&wg, i) // если (*), тогда wg вместо &wg
	}
	wg.Wait()
	fmt.Println("exit")
}
