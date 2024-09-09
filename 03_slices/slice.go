package main

import (
    "fmt"
    "unicode"
    "strconv"
    "sort"
)

func main () {
    usr_inputs := make([]int, 0, 3)

    for {
        var usr_input string
        fmt.Printf("Enter a integer number to add on a list (exit by typing 'x'):\n\t")
        fmt.Scan(&usr_input)

        if unicode.ToLower(rune(usr_input[0])) == rune('x') {
            break
        }

        num_input, err := strconv.Atoi(usr_input)
        if  err != nil {
            continue
        }

        usr_inputs = append(usr_inputs, num_input)

        sort.Ints(usr_inputs)

        fmt.Printf("[")
        for i, value := range usr_inputs {
            if i != 0 {
                fmt.Printf(", ")
            }

            fmt.Printf("%d", value)
        }
        fmt.Printf("]\n")
    }

    fmt.Printf("Final List:\n")
    fmt.Printf("[")
    for i, value := range usr_inputs {
        if i != 0 {
            fmt.Printf(", ")
        }

        fmt.Printf("%d", value)
    }
    fmt.Printf("]\n")

    return
}
