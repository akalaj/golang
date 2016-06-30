package main

import "os"
import "errors"

func main() {
	f := os.Open("filename.ext")
	err :=  errors.New("This is an error message")
	if err != nil {
		log.Fatal(err)
	}
}
// do something with the open *File f
