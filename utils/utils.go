package utils

import (
    "fmt"
    "golang.org/x/net/html"
    "io"
    "log"
    "net/url"
    "path"
    "strings"

    "github.com/disiqueira/gotree"
)

//Given an html.Token, returns an string the value of the href attribute of the token
//If the Token doesn't have any href parameter, it returns false in the second value.
func GetHref(t html.Token) (string, bool) {
    for _, a := range t.Attr {
        if a.Key == "href" {
            return a.Val, true
        }
    }
    return "", false
}

//Given an original path, a base path and the new path, generates a new path
func AppendPath(o string, b string, p string) string {
    var s string
    if path.IsAbs(p) {
        s = o
    } else {
        s = b
    }

    u, err := url.Parse(s)
    if err != nil {
        log.Fatal(err)
    }
    u.Path = path.Join(u.Path, p)
    s = u.String()
    return s
}

//Given an URL it returns a boolean indicating if we can navigate to that URL
func IsValidURL(u string) bool {
    return !strings.Contains(u, ":") && !strings.Contains(u, "//") && !strings.HasPrefix(u, "#")
}

//Given an io.Reader, it looks for <a> attributes and returns it href value in a lis
func ScanLinks(r io.Reader) (ls []string) {
    tn := html.NewTokenizer(r)
    for {
        switch tt := tn.Next(); tt {
        case html.StartTagToken:
            fallthrough
        case html.SelfClosingTagToken:
            t := tn.Token()
            if t.Data == "a" {
                val, ok := GetHref(t)
                if ok && IsValidURL(val) {
                    ls = append(ls, val)
                }
            }
        case html.ErrorToken:
            return ls
        }
    }
}

//Given an url and a tree node, it generates entries to generate a full tree of that url
func generateNode(url string, tree gotree.Tree, urls map[string][]string, or string, printed map[string]bool) {
    urlList := urls[url]
    urlsToPrint := map[string]gotree.Tree{}

    //Doing this separation we prevent trees to be printed in their first occurrence vertically so
    //they are going to be printed in their first occurrence horizontally.
    //Example:
    //We have in the page https://test.com/example/example2 a reference to https://test.com/about
    //Instead of printing the /about tree under /example/example2, we are going to print it in the
    //first level of the tree.
    for _, u := range urlList {
        uNode := tree.Add(u)
        urlBase := AppendPath(or, url, u)
        if !printed[urlBase] {
            printed[urlBase] = true
            urlsToPrint[urlBase] = uNode
        }
    }

    for k, v := range urlsToPrint {
        generateNode(k, v, urls, or, printed)
    }
}

//Given a map from url to list of links in those urls and an origin
//it generates a full tree of the site.
func PrintTree(urls map[string][]string, or string) {
    origin := gotree.New(or)
    printed := map[string]bool{or: true}
    generateNode(or, origin, urls, or, printed)

    fmt.Println(origin.Print())
}
