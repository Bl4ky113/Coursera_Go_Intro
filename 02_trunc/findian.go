package main

import (
    "fmt"
    "strings"
)

const (
    char1 = "i"
    char2 = "a"
    char3 = "n"
)

func main () {
    input := ""
    
    fmt.Printf("Enter a string:\n\t")
    fmt.Scanf("%s", &input)

    switch {
    case !strings.Contains(input, char2):
        fmt.Printf("Not Found!")
        return
    case strings.Index(input, char1) != 0:
        fmt.Printf("Not Found!")
        return
    case strings.LastIndex(input, char3) != len(input) - 1:
        fmt.Printf("Not Found!")
        return
    }

    fmt.Printf("Found!")
}
