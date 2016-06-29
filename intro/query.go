package main

import "database/sql"
import	_ "github.com/go-sql-driver/mysql"
import "os"

func main() {
	db, err := sql.Open("mysql",
		"golang:thisisgolang@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

var (
	id int
	name string
)

rows, err := db.Query("select id, name from users where id = ?", 1)
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
defer rows.Close()
for rows.Next() {
	err := rows.Scan(&id, &name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		
	}
	log.Println(id, name)
}
err = rows.Err()
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
	}
