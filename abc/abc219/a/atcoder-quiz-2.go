package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanf("%d", &x)
	if x < 40 {
		fmt.Println(40 - x)
	} else if x < 70 {
		fmt.Println(70 - x)
	} else if x < 90 {
		fmt.Println(90 - x)
	} else {
		fmt.Println("expert")
	}
}
