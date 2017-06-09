package main 

import (
         "bufio"
         "fmt"
         "os"
         "strings"
 )

func main() {
	consolereader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Month : ")

    input, err := consolereader.ReadString('\n') // this will prompt the user for input

    if err != nil {
            fmt.Println(err)
            os.Exit(1)
         }
    input2 := strings.Replace(input,"\n","",-1)
	
    x := map[string]string{
		"1":"January",
		"2":"February",
		"3":"March"}

	fmt.Printf(x[input2]+"\n")
}

1-25