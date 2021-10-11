package main

import (
	"fmt"
	"sort"
)

func printBool(ok bool) {
	if ok {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func main() {
	var s, t string
	fmt.Scanf("%s %s", &s, &t)

	arr := []string{s, t}
	sort.Strings(arr)

	printBool(arr[0] == s)
}
