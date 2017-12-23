package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("", "temp")
	if nil != err {
		log.Fatalln(err)
	}
	os.RemoveAll(dir)
}
