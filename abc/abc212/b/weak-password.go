package main

import (
	"fmt"
)

func isNext(a, b int) bool {
	if a == 9 {
		return b == 0
	}
	return b-a == 1
}

func main() {
	var x, x1, x2, x3, x4 int
	fmt.Scanf("%d", &x)
	x4 = x % 10
	x3 = (x%100 - x4) / 10
	x2 = (x%1000 - 10*x3 - x4) / 100
	x1 = x / 1000
	ans := "Strong"
	if x1 == x2 && x1 == x3 && x1 == x4 {
		ans = "Weak"
	}
	if isNext(x1, x2) && isNext(x2, x3) && isNext(x3, x4) {
		ans = "Weak"
	}
	fmt.Println(ans)
}
