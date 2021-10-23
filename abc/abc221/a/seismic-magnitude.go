package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	ans := int64(1)
	for i := 0; i < a-b; i++ {
		ans *= 32
	}

	fmt.Println(ans)
}
