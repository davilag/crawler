# WeatherBot
[![CircleCI](https://circleci.com/gh/davilag/crawler/tree/master.svg?style=svg)](https://circleci.com/gh/davilag/crawler/tree/master)
Command to crawl given an URL. 


## Run command
To run the command, you first need to build the project running `go build`. After the `crawler` binary is generated, execute:
```
./crawler http://example.org
```
Which is going to return a tree with the content of the page.