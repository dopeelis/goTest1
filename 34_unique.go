package main

import (
	"fmt"
)

func main() {
	// Задаем строку, которую будем проверять
	str := "heraps12"
	res := isUnique(str)
	// Проверяем вернувшийся результат и выводим ответ в консоль
	if !res {
		fmt.Println("Значения в строке не уникальны")
	} else {
		fmt.Println("Значения в строке уникальны")
	}

}

// // Функция проверки уникальности значения в строке
func isUnique(s string) bool {
	// Создаем мап для помещения всех объектов из строки
	m := make(map[rune]struct{})
	exists := struct{}{}
	// Добавляем руны, остаются только уникальные
	for _, r := range s {
		m[r] = exists
	}
	// Если длина ключей меньше, значит были неуникальные
	if len(m) < len(s) {
		return false
	}

	// Иначе все уникальные
	return true
}
