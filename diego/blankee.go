package main 

import "fmt"
import "strings"

func main() {
	DATE := "1-25-2017"
	fixDATE := strings.Split(DATE, "-")

	fmt.Printf("%v",fixDATE[0:2])
}