package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Создаем функцию, которая выдает рандомные значения
func randNum() int {
	return rand.Intn(100)
}

func main() {
	// Объявляем канал, куда будем записывать значения
	ch := make(chan int, 100)

	// Создаем новый контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			// Если выполнена отмена
			case <-ctx.Done():
				close(ch)
				return
			// Иначе записываем в канал
			default:
				ch <- randNum()
			}

			time.Sleep(500 * time.Millisecond)
		}
	}(ctx)

	// Вызываем отмену через 3 секунды
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	// Выводим значения из канала в консоль
	for i := range ch {
		fmt.Println(i)
	}

	// Обозначаем конец программы
	fmt.Println("finish")
}
