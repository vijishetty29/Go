package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	//Data Source Name Properties
	dsn := mysql.Config{
		User:                 "root",
		Passwd:               "Admin@123",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "go",
		AllowNativePasswords: true,
	}

	//Get db handle
	var err error
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Ping
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	println("Connected %v", db)

	user, err := addUser("Viji", "Shetty", "viji1@gmail.com", "test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Added successfully!", user)
}

func addUser(firstname string, lastname string, email string, password string) (int64, error) {
	println("Connected %v", db)
	result, err := db.Exec("INSERT INTO user (firstname, lastname, email, password) VALUES (?, ?, ?, ?)", firstname, lastname, email, password)
	println("Connected! %v", db)
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addActor: %v", err)
	}
	return id, nil

}
