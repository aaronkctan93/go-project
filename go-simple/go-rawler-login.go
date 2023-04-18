package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

var token string

func main() {
	var url = "https://member.ithome.com.tw/login"
	c := colly.NewCollector()

	// 拿到這次登入的Token
	c.OnHTML("input[name='_token']", func(e *colly.HTMLElement) {
		token = e.Attr("value")
	})

	c.Visit(url)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Chrome/84.0.4147.89 Safari/537.36")
		r.Headers.Set("Host", "https://member.ithome.com.tw")
		r.Headers.Set("Origin", "https://member.ithome.com.tw")
		r.Headers.Set("Referer", "https://member.ithome.com.tw/login")
		// 這幾行在這iT邦幫忙沒有起到作用，但有些網站會照這些資訊判斷、阻擋其他來源
	})

	c.OnScraped(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
		f, err := os.Create("login_token.txt")
		defer f.Close()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err = f.Write([]byte(r.Body))
			checkErr(err)
		}
	})

	var formData = map[string]string{
		"account":  "aaronkctan93",
		"password": "Ithmm620213",
		"_token":   token,
	}
	err := c.Post(url, formData) // 進到該url 執行POST
	if err != nil {
		log.Println(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// https://ithelp.ithome.com.tw/articles/10247492
