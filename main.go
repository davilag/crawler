package main

import (
    "fmt"
    "github.com/davilag/crawler/page"
)

func main() {
    or := "https://twitter.com/"

    //We need a thread safe map here.
    var p page.Page
    p.OriginUrl = or
    p.BaseUrl = b
    e := p.RetrieveUrls()

    fmt.Println(p.Urls)
    if e != nil {
        panic(e)
    }

    // ls := getPageUrls(or, "")
    // for _, val := range ls {
    //     fmt.Println(val)
    // }
}
