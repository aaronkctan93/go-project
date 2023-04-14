package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
	// "html/template"
	// "log"
	// "net/http"
	// "strings"
)

// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm() //解析 url 傳遞的參數，對於 POST 則解析 HTTP 回應內容的主體（request body）
// 	//注意 : 如果沒有呼叫 ParseForm 方法，下面無法取得表單的資料
// 	// fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
// 	// fmt.Println("path", r.URL.Path)
// 	// fmt.Println("scheme", r.URL.Scheme)
// 	// fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("method:", r.Method) //取得請求的方法
// 	if r.Method == "GET" {
// 		t, _ := template.ParseFiles("html/index.html")
// 		log.Println(t.Execute(w, nil))
// 	} else {
// 		//請求的是登入資料，那麼執行登入的邏輯判斷
// 		fmt.Println("username:", r.Form["username"])
// 		fmt.Println("password:", r.Form["password"])
// 		fmt.Println("非GET Method 輸出")
// 	}
// }

func main() {
	// dbtest 為資料庫名稱
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dbtest?charset=utf8")
	checkErr(err)

	// stmt, err := db.Prepare("delete from db_table where username=?")
	// checkErr(err)

	// res, err := stmt.Exec("abstest")
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)
	rows, err := db.Query("SELECT * FROM db_table")
	checkErr(err)

	for rows.Next() {
		var username string
		var password string
		err = rows.Scan(&username, &password)
		checkErr(err)
		fmt.Println(username)
		fmt.Println(password)

		db.Close()
	}
	// http.HandleFunc("/", sayhelloName)       //設定存取的路由
	// http.HandleFunc("/login", login)         //設定存取的路由
	// err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }

	mux := httprouter.New()

	mux.GET("/", index)

	mux.ServeFiles("/css/*filepath", http.Dir("public/css"))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var tmpl = template.Must(template.ParseFiles("html/index.html"))

	tmpl.Execute(w, struct {
		Username string
		Passowrd string
	}{
		"abstest",
		"12345",
	})
}
