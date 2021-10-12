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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()
	s, t := make([]int, n), make([]int, n)
	for i := range s {
		s[i] = nextInt()
	}
	for i := range t {
		t[i] = nextInt()
	}
	for i := 0; i < 2*n; i++ {
		t[(i+1)%n] = min(t[(i+1)%n], t[i%n]+s[i%n])
	}
	for _, ti := range t {
		fmt.Fprintln(wr, ti)
	}
}
