package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scanf("%v %v", &a, &b)
	fmt.Println((a-b)/3 + b)
}
