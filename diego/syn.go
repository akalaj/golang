package main

import "fmt"

type person struct {
	gender string
	age int
	title string
}

func main() {
	human1 := person {
		"Male",
		27,
		"Alex"}
	//Print info
	fmt.Println(human1)
}