package main

import (
    "fmt"
    "os"
    "time"

    "github.com/davilag/crawler/scanner"
)

func main() {
    start := time.Now()

    args := os.Args[1:]

    if len(args) != 1 {
        panic("We need just 1 parameter, the origin url")
    }

    var s scanner.ScannerImp
    urls := scanner.Scan(s, args[0]g)
    end := time.Now()

    fmt.Println("It has taken: ", end.Sub(start))
    fmt.Println("Urls processed: ", len(urls))
}
