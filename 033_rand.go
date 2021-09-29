package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Ставим таймер, чтобы программа не выполнялась вечно
	_ = time.AfterFunc(time.Duration(5)*time.Second, func() {
		fmt.Println("Time is over")
		os.Exit(0)
	})
	// Бесконечным циклом вызываем вывод в консоль
	for {
		i := <-evenChan(randChan())
		fmt.Println(i)
	}
}

// Создаем и записываем в канал рандомные значение
func randChan() chan int {
	outChan := make(chan int)
	go func() {
		for {
			time.Sleep(time.Duration(1) * time.Second)
			outChan <- rand.Int()
		}
	}()
	return outChan
}

// Создаем канал с четными значениями
func evenChan(ch chan int) chan int {
	outChan := make(chan int)
	go func() {
		// Проверяем на четность
		for i := range ch {
			if i%2 == 1 {
				continue
			}
			outChan <- i
		}
	}()
	return outChan
}
