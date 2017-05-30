package main

/*This is a small script I wrote that allows a string to be encapsulated in the same way MySQL does.CLEAR
This is version 1

-WISHLIST-
1.) The ability to accept a heading
*/

import "fmt"
import "bytes"
import "os/exec"

func CLEAR() {
	prod, err := exec.Command("clear").Output()
	if err != nil {
		return
	}
	print(string(prod))
}

func lineGen(LENGTH string) {
	var buffer bytes.Buffer
	for i := 0; i < len(LENGTH)+2; i++ {
		buffer.WriteString("-")
	}
	LID := "+" + buffer.String() + "+"
	VALUE := LID + "\n| " + LENGTH + " |\n" + LID
	fmt.Println(VALUE)
}

func main() {
	CLEAR()
	fmt.Print("Hello. Please give me a string... \n")
	var PHRASE string
	fmt.Scanln(&PHRASE)
	lineGen(PHRASE)
}
