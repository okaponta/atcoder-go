package main

import (
	"fmt"
)

func reverse(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func main() {
	var n int64
	fmt.Scanf("%d", &n)

	answer := make([]rune, 0)

	for n != 0 {
		if n%2 == 0 {
			answer = append(answer, 'B')
			n = n / 2
			continue
		}
		answer = append(answer, 'A')
		n -= 1
	}

	fmt.Println(string(reverse(answer)))
}
