package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("ex.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte("要写入的文本内容"))
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
