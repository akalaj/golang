package main

/*This is an example of how to scan values into a struct*/

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
