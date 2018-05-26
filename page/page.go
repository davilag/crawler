package page

import (
    "fmt"
    "github.com/davilag/crawler/utils"
    "golang.org/x/net/html"
    "net/http"
    "sync"
)

type PageInterface interface {
    RetrieveUrls() error
}

type Page struct {
    BaseUrl   string
    OriginUrl string
    Urls      []string
    UrlMap    sync.Map
    Pages     []Page
}

func (p *Page) RetrieveUrls(o chan bool) {
    done := make(chan bool)
    go p.retrieveUrl(done)
    <-done
    doneChild := make(chan bool)
    for _, url := range p.Urls {
        v := utils.AppendPath(p.OriginUrl, p.BaseUrl, url)
        if _, ok := p.UrlMap.LoadOrStore(v, true); !ok {
            fmt.Println("We haven't checked this url yet")
            fmt.Println(v)
            var np Page
            np.OriginUrl = p.OriginUrl
            np.BaseUrl = v
            np.UrlMap = p.UrlMap
            p.Pages = append(p.Pages, np)
            go np.RetrieveUrls(doneChild)
        }
    }
    for range p.Pages {
        <-doneChild
    }
    o <- true
}

func (p *Page) retrieveUrl(c chan bool) {
    if p.OriginUrl == "" {
        fmt.Println("We haven't received the origin url")
        p.OriginUrl = p.BaseUrl
    }
    var ls []string
    fmt.Println("Retrieving")
    fmt.Println(p.BaseUrl)
    r, e := http.Get(p.BaseUrl)
    if e != nil {
        fmt.Println("We have an error")
        fmt.Println(e)
        c <- false
        return
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
                    ls = append(ls, val)
                }
            }
        case html.ErrorToken:
            fmt.Println(ls)
            p.Urls = ls
            c <- true
        }
    }
}
