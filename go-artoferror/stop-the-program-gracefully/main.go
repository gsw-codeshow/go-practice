package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.OpenFile(`e:\test.txt`, os.O_APPEND, 0644)
	if nil != err {
		fmt.Fprintf(os.Stderr, "Error Opening the file test.txt")
		log.Fatal(err)
		os.Exit(1)
	}
}
