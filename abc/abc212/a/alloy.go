package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	s := "Alloy"
	if a == 0 {
		s = "Silver"
	}
	if b == 0 {
		s = "Gold"
	}
	fmt.Println(s)
}
