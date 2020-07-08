package main

import (
	"database/sql"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	var er error
	db, er := sql.Open("mysql", "root:123456@tcp(localhost:3306)/world?charset=utf8")

	if er != nil {
		panic("db connect error")
	}

	if e := db.Ping(); e != nil {
		panic("ping error")
	}
}

func main() {

}
