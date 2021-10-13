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

func compress(orig []int) map[int]int {
	com := make(map[int]int)
	rank, bef := -1, -1
	for _, e := range orig {
		if e > bef {
			rank++
		}
		com[e] = rank
		bef = e
	}
	return com
}

type Coordinate struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	_, _, n := nextInt(), nextInt(), nextInt()
	a, b, orig := make([]int, n), make([]int, n), make([]Coordinate, n)
	for i := range orig {
		ai, bi := nextInt(), nextInt()
		a[i], b[i], orig[i] = ai, bi, Coordinate{ai, bi}
	}
	sort.Ints(a)
	sort.Ints(b)
	rankA, rankB := compress(a), compress(b)
	for _, c := range orig {
		fmt.Fprintln(wr, rankA[c.x]+1, rankB[c.y]+1)
	}
}
