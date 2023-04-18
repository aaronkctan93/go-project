package main

import (
	"fmt"
)

type user1 struct {
	name string
	Age  int
}

type user2 struct {
	name string
	age  int
	mess string
	// sex  time.Time
}

type User struct {
	u1 user1 //别名
	user2
	Name string
	Age  int
}

func main() {
	var user User
	user.Name = "nick"
	user.u1.Age = 18
	fmt.Println(user) //{{ 18} { 0 {0 0 <nil>}} nick 0}
}
