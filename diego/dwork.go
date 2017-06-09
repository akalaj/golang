package main

import "fmt"
import "time"
import "format"

func main() {
	x := time.Now().Format("Y-m-d H:i:s")
	fmt.Println(x)
}