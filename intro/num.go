package main

 import (
         "bufio"
         "fmt"
         "os"
 )

 func main() {
         consolereader := bufio.NewReader(os.Stdin)
         fmt.Print("Please enter an integer : ")

         input, err := consolereader.ReadString('\n') // this will start the user input

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         fmt.Println(input)
 }
