package main

import (
    "fmt"
    "github.com/davilag/crawler/page"
    "github.com/davilag/crawler/utils"
)

func main() {
    or := "httpg://davidavila.xyz/"

    //We need a thread safe map here.
    var p page.Page
    p.BaseUrl = or
    processing := map[string]bool{or: true}
    processed := map[string]bool{or: true}
    urls := make(map[string][]string)
    // var m sync.Map
    c := make(chan map[string][]string)
    go p.RetrieveUrls(c)
    for len(processing) > 0 {
        out := <-c
        for k, v := range out {
            fmt.Println("storing in urls ", k)
            urls[k] = v
            processed[k] = true
            delete(processing, k)
            for _, u := range v {
                url := utils.AppendPath(or, k, u)
                fmt.Println("Checking in urls ", url)
                if processed[url] || processing[url] {
                    fmt.Println("##################NOPE##################NOPE##################NOPE##################NOPE##################NOPE")
                    continue
                }
                var pv page.Page
                pv.OriginUrl = or
                pv.BaseUrl = url
                processing[url] = true
                go pv.RetrieveUrls(c)
            }
        }
        fmt.Println(processing)
    }

}
