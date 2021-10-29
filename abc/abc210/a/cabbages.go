package main

import (
	"fmt"
)

func main() {
	var n, a, x, y, ans int
	fmt.Scanf("%v %v %v %v", &n, &a, &x, &y)
	if a > n {
		ans = n * x
	} else {
		ans = a*x + (n-a)*y
	}
	fmt.Println(ans)
}
