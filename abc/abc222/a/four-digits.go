package main

import (
	"fmt"
)

func main() {
	var n string
	fmt.Scanf("%s", &n)
	for i := len(n); i < 4; i++ {
		n = "0" + n
	}
	fmt.Println(n)
}
