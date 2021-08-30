package main

import "fmt"

func main() {
	p1 := Action{Human: &Human{"Lola"}, someAct: "run"}
	p1.doSomeAction()
}

type Human struct {
	name string
}

type Action struct {
	*Human
	someAct string
}

func (a *Action) doSomeAction() {
	fmt.Println(a.name, "do some action:", a.someAct)
}
