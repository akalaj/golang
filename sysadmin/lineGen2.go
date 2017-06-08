package main

import "fmt"
import "os/exec"
import "os"
import "bufio"
import "bytes"
import "strings"


func lineGen(HEADER string, VALUES slice[]) {
	// Find larger of 2 strings [STR1 OR STR2]
	if len(STR1) > len(STR2) {
		BUF := len(STR1) + 2
		LESSER := STR2
	} else {
		BUF := len(STR2) + 2
		LESSER := STR1
	}
	//Create buffered string
	WHTSPACE := BUF - len(LESSER)
	BUFSPACE := strings.Repeat(" ", WHTSPACE)
	LESSERWORD := LESSER+BUFSPACE
	DASHES := strings.Repeat("-", BUF)

	//Create lid to seal value
	LID := "+" + DASHES + "+"
	VALUE1 := LID + "\n| " + STR1 + " |\n" + LID
	VALUE2 := VALUE1 + "\n" + LID + ST
}

func CLEAR() {
	prod, err := exec.Command("clear").Output()
	if err != nil {
		return
	}
	print(string(prod))
}

func main() {
	CLEAR()
    READER := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name : ")

    INPUT, err := READER.ReadString('\n') // this will prompt the user for input

    if err != nil {
            fmt.Println(err)
            os.Exit(1)
    }

	RESULT := strings.Replace(INPUT, "\n", "", -1)

    lineGen(RESULT)
}
