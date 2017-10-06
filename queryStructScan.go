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
	//Object that will be populated with animal details
	type animal struct {
		Id   int64
		Name string
	}

	whale := new(animal)

	db, err := sql.Open("mysql", "$USER:$PASS@tcp(akalaj-lab-2003.opslan.smartsheet.com:3306)/$DB")

	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("select id, name from animals where id = ?", 5)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&whale.Id, &whale.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(whale.Name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
