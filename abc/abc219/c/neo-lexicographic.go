package main

import (
	"fmt"
	"sort"
)

func convert(str string, order map[rune]rune) string {
	r_str := []rune(str)
	for i, r := range r_str {
		r_str[i] = order[r]
	}
	return string(r_str)
}

func inverse() {

}

func main() {
	var x string
	fmt.Scanf("%s", &x)
	convertMap := make(map[rune]rune)
	inverseMap := make(map[rune]rune)
	rx := []rune(x)
	for i, r := range rx {
		convertMap[r] = rune('a' + i)
		inverseMap[rune('a'+i)] = r
	}

	var n int
	fmt.Scanf("%d", &n)
	s := make([]string, n)
	for i := range s {
		var si string
		fmt.Scanf("%s", &si)
		s[i] = convert(si, convertMap)
	}

	sort.Strings(s)

	for _, si := range s {
		fmt.Println(convert(si, inverseMap))
	}
}
