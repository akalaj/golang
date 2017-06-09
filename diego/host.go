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

	q := "SHOW GLOBAL VARIABLES LIKE 'hostname'"

	rows, err := db.Query(q)
	if err != nil {
	  log.Fatal(err)
	}

//for rows.Next() {
//	var value string
//	err = rows.Scan(&value)
//	if err != nil {
//	  log.Println(err)
//	}
//	results  := "The hostname for this box is: "+value
	//defer db.close()
	fmt.Println(rows.Scan())
}
	//Close Connection
	//defer db.Close()
//}
