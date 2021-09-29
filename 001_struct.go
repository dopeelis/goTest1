package main

import "fmt"

func main() {
	// Создаем экземпляр струкруты Action
	p1 := Action{Human: Human{"Lola"}, someAct: "run"}

	// Вызываем у него функцию от Action
	p1.doSomeAction()

	// Вызываем у него функцию от Human
	p1.sayHello()
}

// Создаем структуру Human
type Human struct {
	name string
}

// Создаем структуру Action от родительской Human
type Action struct {
	Human
	someAct string
}

// Создаем функцию для Action
func (a *Action) doSomeAction() {
	fmt.Println(a.name, "do some action:", a.someAct)
}

// Создаем функцию для Human
func (h *Human) sayHello() {
	fmt.Println(h.name, "say: Hello!")
}
