package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3, t string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)
	fmt.Scanf("%s", &s3)
	fmt.Scanf("%s", &t)

	answer := ""

	for _, c := range t {
		switch c {
		case '1':
			answer += s1
		case '2':
			answer += s2
		case '3':
			answer += s3
		}
	}

	fmt.Println(answer)
}
