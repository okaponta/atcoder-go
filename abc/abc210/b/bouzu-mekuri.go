package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	deck := []byte(s)
	for i := range deck {
		if deck[i] == '1' {
			if i%2 == 0 {
				fmt.Println("Takahashi")
			} else {
				fmt.Println("Aoki")
			}
			return
		}
	}
}
