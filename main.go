package main

import (
    "github.com/davilag/crawler/page"
    "sync"
)

func main() {
    or := "https://twitter.com/"

    //We need a thread safe map here.
    var p page.Page
    var m sync.Map
    p.BaseUrl = or
    p.UrlMap = m
    done := make(chan bool)
    p.RetrieveUrls(done)

    // ls := getPageUrls(or, "")
    // for _, val := range ls {
    //     fmt.Println(val)
    // }
}
