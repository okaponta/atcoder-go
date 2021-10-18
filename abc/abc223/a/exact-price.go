package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	if x%100 == 0 && x >= 100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
