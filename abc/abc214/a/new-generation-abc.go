package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	if 211 < n {
		fmt.Println(8)
		return
	} else if 125 < n {
		fmt.Println(6)
		return
	}
	fmt.Println(4)
}
