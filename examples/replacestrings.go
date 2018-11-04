package main

import (
    "fmt"
    "strings"
)

func main() {
    value := "Your cat is cute"
    fmt.Println(value)

    // Replace the "cat" with a "dog."
    result := strings.Replace(value, "cat", "dog", -1)
    fmt.Println(result)
}
