package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x, y int
	fmt.Scanf("%d.%d", &x, &y)
	ans := strconv.Itoa(x)
	switch {
	case y <= 2:
		ans += "-"
	case 7 <= y:
		ans += "+"
	}
	fmt.Println(ans)
}
