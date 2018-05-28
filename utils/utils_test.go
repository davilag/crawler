package utils_test

import (
    "strings"
    "testing"

    "github.com/davilag/crawler/utils"
    "github.com/stretchr/testify/assert"
    "golang.org/x/net/html"
)

func TestUtils_GetHref(t *testing.T) {
    scenarios := []struct {
        Token           html.Token
        ExpectedRef     string
        ExpectedSuccess bool
    }{
        {
            Token: html.Token{
                Attr: []html.Attribute{
                    {
                        Key: "href",
                        Val: "value",
                    },
                    {
                        Key: "img",
                        Val: "test.png",
                    },
                },
            },
            ExpectedRef:     "value",
            ExpectedSuccess: true,
        },
        {
            Token: html.Token{
                Attr: []html.Attribute{},
            },
            ExpectedRef:     "",
            ExpectedSuccess: false,
        },
        {
            Token: html.Token{
                Attr: []html.Attribute{
                    {
                        Key: "href",
                        Val: "value",
                    },
                },
            },
            ExpectedRef:     "value",
            ExpectedSuccess: true,
        },
        {
            Token: html.Token{
                Attr: []html.Attribute{
                    {
                        Key: "img",
                        Val: "value.png",
                    },
                },
            },
            ExpectedRef:     "",
            ExpectedSuccess: false,
        },
    }

    for _, s := range scenarios {
        ref, succ := utils.GetHref(s.Token)
        assert.Equal(t, s.ExpectedSuccess, succ)
        assert.Equal(t, s.ExpectedRef, ref)
    }
}

func TestUtils_AppendPath(t *testing.T) {
    scenarios := []struct {
        Origin   string
        Base     string
        Path     string
        Expected string
    }{
        {
            Origin:   "https://test.com/",
            Base:     "https://test.com/",
            Path:     "path",
            Expected: "https://test.com/path",
        },
        {
            Origin:   "https://test.com/",
            Base:     "https://test.com/path",
            Path:     "path2",
            Expected: "https://test.com/path/path2",
        },
        {
            Origin:   "https://test.com/",
            Base:     "https://test.com/test",
            Path:     "/path",
            Expected: "https://test.com/path",
        },
    }

    for _, s := range scenarios {
        result := utils.AppendPath(s.Origin, s.Base, s.Path)
        assert.Equal(t, s.Expected, result)

    }

}

func TestUtils_IsValidURL(t *testing.T) {
    scenarios := []struct {
        Url      string
        Expected bool
    }{
        {
            Url:      "test",
            Expected: true,
        },
        {
            Url:      "/test",
            Expected: true,
        },
        {
            Url:      "#test",
            Expected: false,
        },
        {
            Url:      "https://test.com",
            Expected: false,
        },
        {
            Url:      "test:hello",
            Expected: false,
        },
    }

    for _, s := range scenarios {
        result := utils.IsValidURL(s.Url)
        assert.Equal(t, s.Expected, result)

    }
}

func TestUtils_ScanLinks(t *testing.T) {
    scenarios := []struct {
        Html  string
        Links []string
    }{
        {
            Html: `<html>
                     <a href="link1"></a>
                     <a href="/link2"></a>
                     <a href="https://test.com"></a>
                   <html/>`,
            Links: []string{"link1", "/link2"},
        },
        {
            Html: `<html>
                     <a href="link1"/>
                     <a href="/link2"></a>
                     <a href="https://test.com"></a>
                   <html/>`,
            Links: []string{"link1", "/link2"},
        },
        {
            Html: `<html>
                     <a href="link1:link3"></a>
                     <a href="#link2"></a>
                     <a href="https://test.com"></a>
                   <html/>`,
            Links: []string{},
        },
    }

    for _, s := range scenarios {
        links := utils.ScanLinks(strings.NewReader(s.Html))
        assert.Equal(t, len(s.Links), len(links))
    }
}
