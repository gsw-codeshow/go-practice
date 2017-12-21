package main

import (
	"fmt"
	"net/http"
)

func getLink() (*http.Response, error) {
	content, err := http.Get("www.google.com")
	if nil != err {
		return nil, err
	}
	return content, err
}

func main() {
	resp, err := getLink()
	if nil != err {
		return
	}
	fmt.Println(resp)
}
