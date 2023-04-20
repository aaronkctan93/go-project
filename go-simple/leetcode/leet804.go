package main

import (
	"fmt"
	"strings"
)

func main() {
	morceCode := map[string]string{
		"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".",
		"f": "..-.", "g": "--.", "h": "....", "i": "..", "j": ".---",
		"k": "-.-", "l": ".-..", "m": "--", "n": "-.", "o": "---",
		"p": ".--.", "q": "--.-", "r": ".-.", "s": "...", "t": "-",
		"u": "..-", "v": "...-", "w": ".--", "x": "-..-", "y": "-.--", "z": "--..",
	}

	words := []string{"gin", "zen", "gig", "msg"}
	ans := make([]string, len(words))
	count, compareOK := 0, 0
	merge := []string{}
	for _, word := range words {
		splitWord := strings.Split(word, "")

		for _, checkWord := range splitWord {
			temp, _ := morceCode[checkWord]
			merge = append(merge, temp)

		}
		splitSpace := strings.Join(merge, "")

		ans[count] = splitSpace
		merge = []string{}
		count++

	}
	for i, output := range ans {
		for j := 0; j < len(words); j++ {
			if i != j {
				if ans[j] == output {
					compareOK++
					fmt.Println("j", j, ans[j], " ", output)
				}
			}
		}
	}
	fmt.Println(compareOK)
	// 剩下如何拍出重複計算
}
