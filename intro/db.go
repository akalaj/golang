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
	defer db.Close()
}
