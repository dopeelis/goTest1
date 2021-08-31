package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Пользовательское значение i-го бита для замены
	bitNum := 4

	// Исходное число
	var val int64 = 10234
	fmt.Println("value:", val)

	// Бинарное представление числа
	bitVal := strconv.FormatInt(val, 2)
	// Для конечной проверки
	fmt.Println("binary value:", bitVal)

	// Для изменения бита в порядке "слева направо"
	// Если так не надо, то можно просто убрать данную строку
	// И заменить в сдвиге на bitNum
	realBitNum := len(bitVal) - 1 - bitNum

	// Переключение бита
	newVal := val ^ (1 << realBitNum)
	fmt.Println("new value:", newVal)

	// Проверка, что все заменилось, как нужно
	bitNewVal := strconv.FormatInt(newVal, 2)
	fmt.Println("new binary value:", bitNewVal)
}
