package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	datalist := []int{1, 2, 6, 3, 5, 7, 8, 4, 10}
	fmt.Println(linearsearch(datalist, 7))
}
