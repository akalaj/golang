package main

import "fmt"
import "bytes"

func main() {
	var NUM int = 6
	var buffer bytes.Buffer
	for i := 0; i < NUM; i++ {
    		buffer.WriteString("a")
	}
	fmt.Println(buffer.String())
}