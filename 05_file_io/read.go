package main

import (
    "fmt"
    "log"
    "os"
    "io"
    "bufio"
)

type Name struct {
    fname string
    lname string
}

func main () {
    var filePath string
    var nameArr []Name = make([]Name, 0)

    fmt.Printf("Enter the Path of the File to Read:\t")
    fmt.Scanf("%s", &filePath)

    file, err := os.OpenFile(filePath, os.O_RDONLY, 0o111)
    if err != nil {
        log.Fatal("Error opening the file")
        os.Exit(-1)
    }
    defer file.Close()
    
    fileReader := bufio.NewReader(file)
    for {
        line, _, err := fileReader.ReadLine();
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatalf("Error reading line from file: %s", filePath)
            os.Exit(-1)
        }
        
        var nameObj *Name = new(Name)
        fmt.Sscanf(string(line), "%s %s", &nameObj.fname, &nameObj.lname)
        nameArr = append(nameArr, *nameObj)
    }

    fmt.Printf("List of names collected from %s:\n\n", filePath)
    for i, nameObj := range nameArr {
        fmt.Printf("%d:\t%s %s\n", i + 1, nameObj.fname, nameObj.lname)
    }

    return
}
