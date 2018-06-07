package main

import (
	"fmt"
	"net/http"
)

type Scanner interface {
	FetchLinks(u string, o chan map[string][]string)
}

type ScannerImp struct {
}

//Given an url and a channel, it returns a map from
//the given url to a list of links that it has in the channel
func (s ScannerImp) FetchLinks(u string, o chan map[string][]string) {
	ls, e := fetchPage(u)
	if e != nil {
		fmt.Println("Error retrieving url: ", u)
		fmt.Println(e.Error)
	}
	m := make(map[string][]string)
	m[u] = ls
	o <- m
}

//Given an url, fetchs the html page and returns a list of all
//the links included in that page and an error if ocurred.
func fetchPage(u string) (ls []string, er error) {
	r, e := http.Get(u)
	if e != nil {
		return ls, e
	}

	defer r.Body.Close()
	ls = ScanLinks(r.Body)
	return ls, nil
}

//Method which receives an implementation of Scanner and the origin url.
//It recursively scans all the urls that are found on each page.
func Scan(s Scanner, or string) (urls map[string][]string) {
	//Map which is going to contain the pages that we're processing
	processing := map[string]bool{or: true}
	//Map which is going to contain the pages that we've processed
	processed := map[string]bool{or: true}
	//Map from url to a list of links that the url contains
	urls = make(map[string][]string)

	c := make(chan map[string][]string)
	go s.FetchLinks(or, c)
	for len(processing) > 0 {
		//Channel waiting for the result of the processing of an url
		out := <-c
		for k, v := range out {
			urls[k] = v
			processed[k] = true
			delete(processing, k)
			for _, u := range v {
				url := AppendPath(or, k, u)

				//If we have processed the url or we are processing it,
				//we don't want to fetch it again
				if processed[url] || processing[url] {
					continue
				}
				processing[url] = true
				go s.FetchLinks(url, c)
			}
		}
	}
	return
}
