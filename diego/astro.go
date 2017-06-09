package main 

import (
         "bufio"
         "fmt"
         "os"
 )

func main() {

	consolereader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name : ")

    input, err := consolereader.ReadString('\n') // this will prompt the user for input

    if err != nil {
            fmt.Println(err)
            os.Exit(1)
         }

    fmt.Println(input)
	//x := map[int]string{
	//	1:"January",
	//	2:"February",
	//	3:"March"}

	//fmt.Printf(x[2]+"\n")
}