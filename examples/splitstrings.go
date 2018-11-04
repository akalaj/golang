package main

import (
    "fmt"
    "strings"
)

func main() {
    data := "cat,bird,dog"

    // Split on comma.
    result := strings.Split(data, ",")

    // Display all elements.
    for i := range result {
        fmt.Println(result[i])
    }
    // Length is 3.
    fmt.Println(len(result))
}
