package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Импортируем библиотеку для работы с большими числами
	// и объявляем наши числа через эту библиотеку
	// Float необходим для корректного вывода значения деления
	a := big.NewFloat(1500000)
	b := big.NewFloat(1300000)

	// Производим операции сложения, вычитания, деления и умножения
	sum := new(big.Float).Add(a, b)

	sub := new(big.Float).Sub(a, b)

	div := new(big.Float).Quo(a, b)

	mul := new(big.Float).Mul(a, b)

	// Выводим в консоль
	fmt.Printf("Sum: %1.f\n", sum) // Для вывода в читаемом виде
	fmt.Println("Sub:", sub)
	fmt.Println("Div:", div)
	fmt.Printf("Mul: %1.f\n", mul) // Для вывода в читаемом виде
}
