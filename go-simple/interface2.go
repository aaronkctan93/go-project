package main

import (
	"fmt"
)

type Student struct {
	Name string
}

func judgement(items ...interface{}) {
	for k, v := range items {
		switch v.(type) {
		case string:
			fmt.Printf("string, %d[%v]\n", k, v)
		case int, int32, int64:
			fmt.Printf("int, %d[%v]\n", k, v)

		case float32, float64:
			fmt.Printf("float, %d[%v]\n", k, v)

		case Student:
			fmt.Printf("Student, %d[%v]\n", k, v)
		case *Student:
			fmt.Printf("Student, %d[%p]\n", k, v)

		}
	}
}

func main() {

	stu1 := Student{Name: "Alan"}
	judgement(1, 1.2, "str", stu1)

}
