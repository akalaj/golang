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
	}

	q := "SELECT id, name FROM testing.animals WHERE id = ?"

	rows, err := db.Query(q, 1)
	if err != nil {
	  log.Fatal(err)
	}

for rows.Next() {
	var id int
	var name string
	err = rows.Scan(&id, &name)
	if err != nil {
	  log.Println(err)
	}
	result1 := id
	result2 := name
	fmt.Println(result1,":",result2)
}
	//Close Connection
	defer db.Close()
}
