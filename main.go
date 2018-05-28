package main

import (
    "fmt"
    "os"

    "github.com/davilag/crawler/scanner"
    "github.com/davilag/crawler/utils"
)

func main() {

    args := os.Args[1:]

    if len(args) != 1 {
        panic("We need just 1 parameter, the origin url")
    }
    or := args[0]
    var s scanner.ScannerImp
    fmt.Println("Scanning url...")
    urls := scanner.Scan(s, or)

    utils.PrintTree(urls, or)
}
