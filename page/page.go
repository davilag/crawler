package page

import (
    "github.com/davilag/crawler/utils"
    "golang.org/x/net/html"
    "net/http"
)

type (
    Page struct {
        BaseUrl   string
        OriginUrl string
        Urls      []string
    }
)

func (p *Page) RetrieveUrls() error {
    return p.retrieveUrl()
}

func (p *Page) retrieveUrl() error {
    var ls []string
    r, e := http.Get(p.BaseUrl)
    if e != nil {
        return e
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
            p.Urls = ls
            return nil
        }
    }
}
