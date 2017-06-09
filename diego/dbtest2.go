package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql",
		"akalaj:kalaj@tcp(127.0.0.1:3306)/testing")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Database connection successful")
	}
	defer db.Close()
}
