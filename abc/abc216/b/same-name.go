package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	names := make([]string, n)
	for i := range names {
		var s, t string
		fmt.Scanf("%s %s", &s, &t)
		names[i] = s + "_" + t
	}
	sort.Strings(names)
	for i := 0; i < n-1; i++ {
		if names[i] == names[i+1] {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
