package main

import (
    "fmt"
    "unicode"
    "strconv"
    "sort"
)

func main () {
    usr_inputs := make([]int, 3)
    current_input := 0

    for {
        var usr_input string
        fmt.Printf("Enter a integer number to add on a list (exit by typing 'x'):\n\t")
        fmt.Scan(&usr_input)

        if unicode.ToLower(rune(usr_input[0])) == rune('x') {
            break
        }

        if num_input, err := strconv.Atoi(usr_input); err == nil {
            if current_input < len(usr_inputs) {
                usr_inputs[current_input] = num_input
            } else {
                usr_inputs = append(usr_inputs, num_input)
            }

            current_input++
        }
    }

    sort.Ints(usr_inputs)

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
