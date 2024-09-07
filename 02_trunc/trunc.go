package main

import "fmt"

func main () {
    var usr_input float64
    var int_input int

    fmt.Printf("Enter a Number with decimal part:\n")
    fmt.Scanf("%f", &usr_input)
    int_input = int(usr_input)
    fmt.Printf("The Integer part of the number is: %d\n", int_input)
}
