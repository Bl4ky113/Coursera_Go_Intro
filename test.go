package main

import "fmt"

func main () {
    arr := [...]int{1, 2, 3}
    sli := arr[0:2]
    fmt.Printf("%d\t%d\n%d\t%d", len(arr), cap(arr), len(sli), cap(sli))
}
