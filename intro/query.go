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
	var id int
	var str string
	q := "SELECT id, animal, name FROM go.animals WHERE id = ?"

	rows, err := db.Query(q, 1)
	if err != nil {
	  log.Fatal(err)
	}
	//Close Connection
	defer db.Close()
}
