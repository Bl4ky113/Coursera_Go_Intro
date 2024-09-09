package main

import (
    "fmt"
    "os"
    "encoding/json"
)

func main () {
    var usrName string = ""
    var usrAddress string = ""
    var usrMap map[string]string = make(map[string]string)
    
    fmt.Printf("Enter your Name:\t")
    if _, err := fmt.Scanf("%s", &usrName); err != nil {
        fmt.Printf("Error getting user input\n")
        os.Exit(-1)
    }

    fmt.Printf("Enter your Address:\t")
    if _, err := fmt.Scanf("%s", &usrAddress); err != nil {
        fmt.Printf("Error getting user input\n")
        os.Exit(-1)
    }

    usrMap["name"] = usrName
    usrMap["address"] = usrAddress
    jsonStr, err := json.Marshal(usrMap)
    if err != nil {
        fmt.Printf("Error parsing Map into JSON")
        os.Exit(-1)
    }
    fmt.Printf("%s\n", jsonStr)

    return
}
