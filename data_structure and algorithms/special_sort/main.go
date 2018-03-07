package main

import (
	"fmt"
	"math/rand"
)

func InitStartList() (start [100]int) {
	var flag [100]bool
	insert_num := int(0)
	for {
		i := rand.Int() % 100
		if !flag[i] {
			start[insert_num] = i
			insert_num++
			flag[i] = true
		}
		if 100 == insert_num {
			break
		}
	}
	return
}

func sort(start *[100]int) [100]int {
	for i := 0; i < 100; i++ {
		for {
			if i != start[i] {
				start[i], start[start[i]] = start[start[i]], start[i]
				continue
			}
			break
		}
	}
	return *start
}

func main() {
	start := InitStartList()
	sort(&start)
	fmt.Println(start)
}
