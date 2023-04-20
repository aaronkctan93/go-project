package main

import (
	"fmt"
)

func test1(i interface{}) {
	n, ok := i.(int)
	if !ok {
		fmt.Println("error")
	}

	n += 10
	fmt.Println(n)
}

func main() {
	test1("OK")
}
