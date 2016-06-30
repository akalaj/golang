package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
	"golang:thisisgolang@tcp(127.0.0.1:3306)/go")
	if err != nil {
	  log.Println(err)
	}
	var str string
	err = db.QueryRow(
	  "SELECT name FROM go.animals WHERE id=1").Scan(&str)
	if err != nil && err != sql.ErrNoRows {
	  log.Println(err)
	}
	  log.Println(str)
	defer db.Close()
}
