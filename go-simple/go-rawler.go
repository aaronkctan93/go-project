package main

import (
	"fmt"
	"os"
	// _ "fmt"

	"github.com/gocolly/colly"
)

func main() {
	f, err := os.Create("ex.txt")
	defer f.Close()

	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write([]byte(r.Body))
			checkErr(err)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
