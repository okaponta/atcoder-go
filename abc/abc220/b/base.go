package main

import (
	"fmt"
	"strconv"
)

func main() {
	var k int
	var a, b string
	fmt.Scanf("%d", &k)
	fmt.Scanf("%s %s", &a, &b)

	a10, _ := strconv.ParseInt(a, k, 0)
	b10, _ := strconv.ParseInt(b, k, 0)
	fmt.Printf("%d\n", int64(a10*b10))
}
