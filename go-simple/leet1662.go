package main

import (
	"fmt"
)

func main() {
	word1 := []string{"abc", "d", "defg"}
	// word2 := []string{"abcddefg"}
	// var isTrue bool
	strConvert1 := make([]string, 0)
	// strConvert1 := strings.Join(word1, "")
	// strConvert2 := strings.Join(word2, "")

	for _, value := range word1 {
		strConvert1 = append(strConvert1, value)
	}

	// if strConvert1 == strConvert2 {
	// 	isTrue = true
	// } else {
	// 	isTrue = false
	// }

	fmt.Println(strConvert1)
	// fmt.Println(strConvert2)
	// fmt.Println(isTrue)
}
