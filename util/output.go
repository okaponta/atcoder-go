package util

import (
	"fmt"
	"strings"
)

func printBool(ok bool) {
	if ok {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func printSlice(slice []string) {
	joined := strings.Join(slice, " ")
	fmt.Println(joined)
}
