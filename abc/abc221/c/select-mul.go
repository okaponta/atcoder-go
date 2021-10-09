package main

import (
	"fmt"
	"sort"
	"strconv"
)

func testFirst(first, second []rune, r rune) int {
	first = append(first, r)
	firstInt, _ := strconv.Atoi(string(first))
	secondInt, _ := strconv.Atoi(string(second))
	return firstInt * secondInt
}

func testSecond(first, second []rune, r rune) int {
	second = append(second, r)
	firstInt, _ := strconv.Atoi(string(first))
	secondInt, _ := strconv.Atoi(string(second))
	return firstInt * secondInt
}

func main() {
	var n string
	fmt.Scanf("%s", &n)

	runes := []rune(n)
	sort.Slice(runes, func(i, j int) bool { return runes[i] > runes[j] })

	first := make([]rune, 0)
	second := make([]rune, 0)

	for i, j := range runes {
		if i == 0 {
			first = append(first, j)
			continue
		}
		if i == 1 {
			second = append(second, j)
			continue
		}
		ansFirst := testFirst(first, second, j)
		ansSecond := testSecond(first, second, j)
		addToFirst := ansFirst > ansSecond
		if addToFirst {
			first = append(first, j)
		} else {
			second = append(second, j)
		}
	}

	firstInt, _ := strconv.Atoi(string(first))
	secondInt, _ := strconv.Atoi(string(second))

	fmt.Println(firstInt * secondInt)
}
