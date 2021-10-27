package main

import (
	"fmt"
)

const mod = 1000000007

func addCount(i int, count []int) int {
	return (count[i-1] + count[i]) % mod
}

func main() {
	var s string
	fmt.Scanf("%s", &s)
	bytes := []byte(s)
	counts := make([]int, 8)

	for _, b := range bytes {
		switch b {
		case 'c':
			counts[0]++
		case 'h':
			counts[1] = addCount(1, counts)
		case 'o':
			counts[2] = addCount(2, counts)
		case 'k':
			counts[3] = addCount(3, counts)
		case 'u':
			counts[4] = addCount(4, counts)
		case 'd':
			counts[5] = addCount(5, counts)
		case 'a':
			counts[6] = addCount(6, counts)
		case 'i':
			counts[7] = addCount(7, counts)
		}
	}
	fmt.Println(counts[7])
}
