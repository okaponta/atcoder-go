package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a == c {
		fmt.Println(a)
		return
	}
	ans := a - a%c + c
	if ans <= b {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}
