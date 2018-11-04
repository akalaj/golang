package main

import (
//	_ "github.com/go-sql-driver/mysql"
//	"database/sql"
	"bufio"
	"os"
	"fmt"
	"log"
	"strings"
)

func main() {

	//Collect horoscope data from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello Please provide your birthday!: ")
	sdate, err := reader.ReadString('\n')
	if err != nil {
	  log.Println(err)
	}

	sdate = strings.Replace(sdate, "\n", "", -1)
	splitdate := strings.Split(sdate, "/")

	fmt.Println(splitdate)
	// 2018-12-18

	//fmt.Sprintf("2018-%s-%s", splitdate[0])

	//fmt.Println(splitdate)

	//Connect to database and work
//	db, err := sql.Open("mysql","go:lang@tcp(127.0.0.1:3306)/horoscope")
//	if err != nil {
//	  log.Println(err)
//	}
//
//	defer db.Close()
//
//	var str string
//
//	err = db.QueryRow("SELECT sign FROM horoscope.signs WHERE id=1").Scan(&str)
//	if err != nil && err != sql.ErrNoRows {
//	  log.Println(err)
//	}
//
//	fmt.Println(str)
}