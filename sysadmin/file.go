package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	//Connect to the database
    db, err := sql.Open("mysql", "script:@sqlpass(127.0.0.1:3306)/radar")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    //Test insert into radar.hostnames
    insertDB, err := db.Exec(
        "INSERT INTO radar.hosts (hostname, service) VALUES('hostname1', 'mysql')")
    if err != nil {
        log.Fatal(err)
    }

    //Return if write was successful
    rowCount, err := insertDB.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("inserted %d rows", rowCount)

    //Read newly inserted row into rows object
    rows, err := db.Query("SELECT * FROM radar.hosts")
    if err != nil {
        log.Fatal(err)
    }

    //For rows found within rows object, scan into the "s" variable and 
    for rows.Next() {
        var one string
        var two string
        var three string
        err = rows.Scan(&s)
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("found row containing %q", s)
    }
    rows.Close()
}
