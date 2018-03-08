package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	var pre, next, sum int
	pre = 0
	next = 1
	count := -1
	return func() int {
		count++
		if count < 2 {
			return count
		}
		sum = pre + next
		pre = next
		next = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
