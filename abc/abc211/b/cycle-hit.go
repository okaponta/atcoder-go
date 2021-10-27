package main

import (
	"fmt"
)

func main() {
	var s1, s2, s3, s4 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)
	fmt.Scanf("%s", &s3)
	fmt.Scanf("%s", &s4)
	hits := make(map[string]struct{})
	hits[s1] = struct{}{}
	hits[s2] = struct{}{}
	hits[s3] = struct{}{}
	hits[s4] = struct{}{}
	_, ok1 := hits["H"]
	_, ok2 := hits["2B"]
	_, ok3 := hits["3B"]
	_, ok4 := hits["HR"]
	if ok1 && ok2 && ok3 && ok4 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
