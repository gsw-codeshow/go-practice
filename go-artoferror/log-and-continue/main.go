package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("baidu.com")
	if nil != err {
		log.Println(err)
	}
	_, err = http.Get("google.com")
	log.Println(err)
}
