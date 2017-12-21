package main

import (
	"fmt"
	"net/http"
	"time"
)

func waitForResponse() error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	url := "www.google.com"
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		fmt.Println("Server not responding. Retrying again")
	}
	return fmt.Errorf("Server faild to respond after the time")
}

func main() {
	err := waitForResponse()
	if nil != err {
		return
	}
}
