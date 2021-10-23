package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)
	if s == t {
		fmt.Println("Yes")
		return
	}

	ss := []byte(s)

	for i := 0; i < len(ss)-1; i++ {
		ss[i], ss[i+1] = ss[i+1], ss[i]
		if string(ss) == t {
			fmt.Println("Yes")
			return
		}
		ss[i], ss[i+1] = ss[i+1], ss[i]
	}
	fmt.Println("No")
}
