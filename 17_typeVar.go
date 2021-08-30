package main

import "fmt"

func main() {
	// Исхрдные данные
	data1 := "hello world"
	data2 := 345
	data3 := true
	data4 := make(chan int)

	// Вывод типа переменной в консоль
	fmt.Println("1:", dataType(data1))
	fmt.Println("2:", dataType(data2))
	fmt.Println("3:", dataType(data3))
	fmt.Println("4:", dataType(data4))
}

// Функция проверки типа переменной
func dataType(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case bool:
		return "bool"
	default:
		_ = t
		return "unknown type"
	}
}
