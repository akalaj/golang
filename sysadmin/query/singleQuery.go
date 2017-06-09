package main

/*This is an example of how to retrieve a single value from a MySQL database. 
The values in the connection string on line 14 have been converted into BASH variables
and are intended to be modified by the user for functionality*/

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql","$USER:$PASSWORD@tcp(127.0.0.1:3306)/$DATABASE")
	
	if err != nil {
		log.Println(err)
	}

	var str string
	
	err = db.QueryRow("SELECT name FROM test.animals WHERE id=1").Scan(&str)
	
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	  fmt.Println(str)
	defer db.Close()
}