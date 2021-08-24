package main

import (
	"fmt"
	"math"
)

// Описываем структуру пары чисел
type pairNum struct {
	a int
	b int
	// Параметры удовлетворения условию
	aCondition bool
	bCondition bool
}

// Функция проверки условия
func (p *pairNum) checkCondition() {
	if p.a > int(math.Pow(2, 20)) {
		p.aCondition = true
	}
	if p.b > int(math.Pow(2, 20)) {
		p.bCondition = true
	}
}

func main() {
	// Задаем пару числе
	pair1 := pairNum{
		a: 56, b: 12}
	// Проверяем условие
	pair1.checkCondition()
	fmt.Println(pair1)
	sum := sum(pair1)
	sub := sub(pair1)
	div := div(pair1)
	mul := mul(pair1)
	// Выводим результаты операций в консоль
	fmt.Printf("Sum: %d\nSub: %d\nDiv: %f\nMul: %d\n", sum, sub, div, mul)
}

// Функция расчета суммы чисел
func sum(p pairNum) int {
	if !p.aCondition && !p.bCondition {
		panic("The condition was not met")
	}
	return p.a + p.b
}

// Функция расчета разности чисел
func sub(p pairNum) int {
	if !p.aCondition && !p.bCondition {
		panic("The condition was not met")
	}
	return p.a - p.b
}

// Функция расчета частного чисел
func div(p pairNum) float64 {
	if !p.aCondition && !p.bCondition {
		panic("The condition was not met")
	}
	return float64(p.a) / float64(p.b)
}

// Функция расчета произведения чисел
func mul(p pairNum) int {
	if !p.aCondition && !p.bCondition {
		panic("The condition was not met")
	}
	return (p.a * p.b)
}
