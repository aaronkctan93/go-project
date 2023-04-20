package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// var url = "https://member.ithome.com.tw/login"
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("2")
	})

	c.OnResponseHeaders(func(r *colly.Response) {
		fmt.Println("3")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))

	})

	c.OnHTML(".qa-list__title-link", func(e *colly.HTMLElement) {
		fmt.Println(e.Text, e.Attr("href"))
	})
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.OnXML("//footer", func(e *colly.XMLElement) {
		fmt.Println("6")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("7")
	})

	// c.Visit(url)
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=1")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=2")
	c.Visit("https://ithelp.ithome.com.tw/users/20125192/ironman/3155?page=3")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// https://ithelp.ithome.com.tw/articles/10247126
