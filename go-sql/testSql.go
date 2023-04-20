package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "root"
	password = "Root1234"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test_database"
)

var DB *sql.DB

type Account struct {
	username string
	password string
}

func InitDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ")/", dbName, "?charset=utf8"}, "")
	DB, _ := sql.Open("mysql", path)
	//設置數據庫最大連接數
	DB.SetConnMaxLifetime(100)
	//設置上數據庫最大閒置連接數
	DB.SetMaxIdleConns(10)
	//驗證連接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
}

func InsertData(acc Account) bool {
	//开启事务
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}
	//准备sql语句
	stmt, err := tx.Prepare("INSERT INTO account (`username`, `password`) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	//将参数传递到sql语句中并且执行
	res, err := stmt.Exec(acc.username, acc.password)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}
	//将事务提交
	tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	return true
}

func QueryData() {
	var acc Account
	rows, e := DB.Query("select * from account")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		e := rows.Scan(acc.username, acc.password)

		if e != nil {
			fmt.Println(json.Marshal(acc))
		}
	}
	rows.Close()
}

func main() {
	// acc := &Account{
	// 	username: "Aaron",
	// 	password: "Qwe123",
	// }
	InitDB()
	// InsertData(*acc)
	QueryData()
	defer DB.Close()
}
