package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n, p := nextInt(), nextInt()
	count := 0
	for i := 0; i < n; i++ {
		if nextInt() < p {
			count++
		}
	}
	fmt.Println(count)
}
