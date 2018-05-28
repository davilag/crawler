package main

import (
    "reflect"
    "testing"

    "github.com/stretchr/testify/assert"
)

var urlsTest = map[string][]string{
    "https://test.com": []string{
        "/about",
        "test",
    },
    "https://test.com/about": []string{
        "/hello",
        "about2",
    },
    "https://test.com/test": []string{
        "test2",
        "/hello",
    },
    "https://test.com/about/about2": []string{
        "about3",
    },
    "https://test.com/hello":               []string{},
    "https://test.com/about/about2/about3": []string{},
    "https://test.com/test/test2":          []string{},
}

type ScannerTest struct {
}

func (s ScannerTest) FetchLinks(u string, o chan map[string][]string) {
    out := map[string][]string{}
    out[u] = urlsTest[u]
    o <- out
}

func TestScanner_Scan(t *testing.T) {
    var s ScannerTest
    urls := Scan(s, "https://test.com")
    assert.Equal(t, len(urls), len(urlsTest))
    for k, _ := range urls {
        assert.True(t, reflect.DeepEqual(urls[k], urlsTest[k]))
    }
}
