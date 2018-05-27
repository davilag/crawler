package page

import (
    "fmt"
    "github.com/davilag/crawler/utils"
    "golang.org/x/net/html"
    "net/http"
)

type PageInterface interface {
    RetrieveUrls() error
}

type Page struct {
    BaseUrl   string
    OriginUrl string
    // UrlMap    sync.Map
}

func (p *Page) RetrieveUrls(o chan map[string][]string) {
    ls, e := p.retrieveUrl()
    if e != nil {
        panic(e)
    }
    m := make(map[string][]string)
    fmt.Println(ls)
    m[p.BaseUrl] = ls
    o <- m
}

func (p *Page) retrieveUrl() (ls []string, er error) {
    if p.OriginUrl == "" {
        fmt.Println("We haven't received the origin url")
        p.OriginUrl = p.BaseUrl
    }

    fmt.Println("Retrieving ", p.BaseUrl)
    r, e := http.Get(p.BaseUrl)
    if e != nil {
        fmt.Println("We have an error")
        fmt.Println(e)
        return ls, e
    }

    defer r.Body.Close()
    tn := html.NewTokenizer(r.Body)
    for {
        switch tt := tn.Next(); tt {
        case html.StartTagToken:
            t := tn.Token()
            if t.Data == "a" {
                val, ok := utils.GetHref(t)
                if ok && utils.IsValidURL(val) {
                    fmt.Println(val)
                    ls = append(ls, val)
                }
            }
        case html.ErrorToken:
            return ls, nil
        }
    }
}
