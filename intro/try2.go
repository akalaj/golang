package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	db, err := sql.Open("mysql",
		"golang:thisisgolang@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer db.Close()
}
