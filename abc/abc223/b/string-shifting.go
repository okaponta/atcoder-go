package main

import (
	"fmt"
	"sort"
)

func shift(s string) string {
	b := []byte(s)
	return string(append(b[1:], b[0]))
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	n := len(s)
	shifts := make([]string, n)
	shifts[0] = s
	for i := 1; i < n; i++ {
		shifts[i] = shift(shifts[i-1])
	}
	sort.Strings(shifts)
	fmt.Println(shifts[0])
	fmt.Println(shifts[n-1])
}
