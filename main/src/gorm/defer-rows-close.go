package main

import (
	db "database/sql"
	"fmt"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// maxConn > 1
	fewConns()
	// maxConn = 1
	oneConn()

	// maxConn = 1 + for rows.Next()
	oneConnWithRowsNext()
	// maxConn = 1 + for rows.Next() + 提前退出
	oneConnWithRowsNextWithError()
	// maxConn = 1 + for rows.Next() + 提前退出 + defer rows.Close()
	oneConnWithRowsNextWithErrorWithRowsClose()
}

func fewConns() {
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(10)
	rows, err := db.Query("select * from test where name = 'jackie' limit 10")
	if err != nil {
		fmt.Println("query error")
	}

	row, _ := db.Query("select * from test")
	fmt.Println(row, rows)
}

func oneConn() {
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(1)
	rows, err := db.Query("select * from test where name = 'jackie' limit 10")
	if err != nil {
		fmt.Println("query error")
	}

	row, _ := db.Query("select * from test")
	fmt.Println(row, rows)
}

func oneConnWithRowsNext() {
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(1)
	rows, err := db.Query("select * from test where name = 'jackie' limit 10")
	if err != nil {
		fmt.Println("query error")
	}

	for rows.Next() {
		fmt.Println("close")
	}

	row, _ := db.Query("select * from test")
	fmt.Println(row, rows)
}

func oneConnWithRowsNextWithError() {
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(1)
	rows, err := db.Query("select * from test where name = 'jackie' limit 10")
	if err != nil {
		fmt.Println("query error")
	}

	i := 1
	for rows.Next() {
		i++
		if i == 3 {
			break
		}
		fmt.Println("close")
	}

	row, _ := db.Query("select * from test")
	fmt.Println(row, rows)
}

func oneConnWithRowsNextWithErrorWithRowsClose() {
	db, _ := db.Open("mysql", "root:rootroot@/dqm?charset=utf8&parseTime=True&loc=Local")

	db.SetMaxOpenConns(1)
	rows, err := db.Query("select * from test where name = 'jackie' limit 10")
	if err != nil {
		fmt.Println("query error")
	}

	i := 1
	for rows.Next() {
		i++
		if i == 3 {
			break
		}
		fmt.Println("close")
	}
	rows.Close()

	row, _ := db.Query("select * from test")
	fmt.Println(row, rows)
}
