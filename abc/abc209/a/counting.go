package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d", &a, &b)
	if a > b {
		c = 0
	} else {
		c = b - a + 1
	}
	fmt.Println(c)
}
