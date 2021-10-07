package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scanf("%d", &n)
	count := 0

	for n > 1 {
		n = n >> 1
		count++
	}
	fmt.Println(count)
}
