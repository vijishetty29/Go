package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Admin@123@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO testtable VALUES('25')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Added successfully!")
}
