package main

import (
	"fmt"
)

type Human interface {
	Walk()
	Learn(n string)
}

type men struct {
	Name string
}

func (x men) Walk() {
	fmt.Println("Xuan Walked.")
}

func (x men) Learn(n string) {
	fmt.Printf("Xuan is learning the %s book.\n", n)
}

func main() {
	var m men

	m.Name = "Xuan"
	m.Walk()
	m.Learn("Golang")
}
