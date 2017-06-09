package main

/*This is an example of how to retrieve multiple values from a single MySQL database record.
The values in the connection string on line 14 have been converted into BASH variables
and are intended to be modified by the user for functionality*/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "$USER:$PASSWORD@tcp(127.0.0.1:3306)/$DATABASE")

	if err != nil {
		log.Println(err)
	}

	var id int
	var name string

	rows, err := db.Query("select id, name from animals where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
