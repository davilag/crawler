package main

import (
    "fmt"
    "os"
)

func main() {

    args := os.Args[1:]

    if len(args) != 1 {
        panic("We need just 1 parameter, the origin url")
    }
    or := args[0]
    var s ScannerImp
    fmt.Println("Scanning url...")
    urls := Scan(s, or)

    PrintTree(urls, or)
}
