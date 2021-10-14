package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextTown(start, before int, route [][]int) {
	fmt.Fprint(wr, start, " ")
	for _, next := range route[start] {
		if next != before {
			nextTown(next, start, route)
			fmt.Fprint(wr, start, " ")
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n := nextInt()
	route := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		route[a] = append(route[a], b)
		route[b] = append(route[b], a)
	}

	for i := 1; i <= n; i++ {
		sort.Ints(route[i])
	}

	nextTown(1, 0, route)
	fmt.Fprintln(wr)
}
