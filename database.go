package main

import(
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_go")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error on ping: ", err)
	}
	fmt.Println("Connected to database")
}