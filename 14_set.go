package main

import "fmt"

var exists = struct{}{}

// Создаем структуру, в которой будут наши значения
type set struct {
	m map[string]struct{}
}

// Добвляем нашей структуре функцию создания
func newSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

// И функцию добавления элемента
func (s *set) add(value string) {
	s.m[value] = exists
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Исходный массив:", arr)

	set := newSet()
	for _, i := range arr {
		set.add(i)
	}

	fmt.Println("Set:", set)
}
