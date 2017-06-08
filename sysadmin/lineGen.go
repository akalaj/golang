package main

import "fmt"
import "os/exec"
import "os"
import "bufio"
import "bytes"
import "strings"

/*This is a small script I wrote that allows a string to be encapsulated in the same way MySQL does.CLEAR
This is version 1
-WISHLIST-
1.) The ability to terminate based on enter rather than '\n' */

func lineGen(LENGTH string) {
	var buffer bytes.Buffer
	for i := 0; i < len(LENGTH)+2; i++ {
		buffer.WriteString("-")
	}
	LID := "+" + buffer.String() + "+"
	VALUE := LID + "\n| " + LENGTH + " |\n" + LID
	fmt.Println(VALUE)
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
