package main

import "fmt"

func main() {
	var s, t int
	fmt.Scanf("%d %d", &s, &t)
	count := 0
	for i := 0; i <= s; i++ {
		for a := 0; a <= i; a++ {
			for b := 0; a+b <= i; b++ {
				c := i - a - b
				if a*b*c <= t {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
