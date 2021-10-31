package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n, x := nextInt(), nextInt()
	sum := 0
	for i := 0; i < n; i++ {
		sum += nextInt()
		if i%2 != 0 {
			sum--
		}
	}
	if sum <= x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
