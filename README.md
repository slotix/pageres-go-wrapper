Pageres Go Wrapper
----------
[![Build Status](https://travis-ci.org/slotix/slugifyurl.svg?branch=master)](https://travis-ci.org/slotix/slugifyurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/slotix/slugifyurl)](https://goreportcard.com/report/github.com/slotix/slugifyurl)

Golang package for capturing screenshots of websites in various resolutions. 
It uses Pageres https://github.com/sindresorhus/pageres-cli internally.  

Installation 
------------

```
$ npm install --global pageres-cli
$ go get -u github.com/slotix/pageres-go-wrapper
```

Usage
-----

```go
package main
import (
    "fmt"
    "os"
	sshot "github.com/slotix/pageres-go-wrapper"
)

func main() {
    shotsDir := "shots"
    os.Mkdir(shotsDir, 0777)
	params := Parameters{
		Command: "pageres",
		Sizes:   "1024x768",
		Crop:    "--crop",
		Scale:   "--scale 0.9",
		Timeout: "--timeout 30",
		Filename:  fmt.Sprintf("--filename=%s/<%%= url %%>", shotsDir),
		UserAgent: "",
	}
	urls := []string{
		"http://google.com",
		"https://dbconvert.com",
		"http://something-that-doesnot-exists.com",
	}
	sshot.GetShots(urls, params)
	sshot.DeleteZeroLengthFiles(shotsDir)
}
```
Find more information about parameters at https://github.com/sindresorhus/pageres