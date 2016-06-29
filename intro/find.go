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

         if input = 10 {
	fmt.Println("This number is 10")
	} else {
	fmt.Println("This numbah ain't 10 motha fuckah!")
	}
 }
