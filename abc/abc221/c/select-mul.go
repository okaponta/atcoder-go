package main

import (
	"fmt"
	"sort"
	"strconv"
)

func testFirst(first, second []byte, r byte) int {
	first = append(first, r)
	return multiply(first, second)
}

func testSecond(first, second []byte, r byte) int {
	second = append(second, r)
	return multiply(first, second)
}

func multiply(a, b []byte) int {
	first, _ := strconv.Atoi(string(a))
	second, _ := strconv.Atoi(string(b))
	return first * second
}

func main() {
	var n string
	fmt.Scanf("%s", &n)

	digit := []byte(n)
	sort.Slice(digit, func(i, j int) bool { return digit[i] > digit[j] })

	first := []byte{digit[0]}
	second := []byte{digit[1]}

	for i := 2; i < len(digit); i++ {
		j := digit[i]
		if testFirst(first, second, j) > testSecond(first, second, j) {
			first = append(first, j)
		} else {
			second = append(second, j)
		}
	}

	fmt.Println(multiply(first, second))
}
