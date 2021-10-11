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

type Node struct {
	cut  int
	next *Node
}

type CutList struct {
	root *Node
	len  int
}

type Query struct {
	c, x int
}

func print(cut []int, x, l int) {
	low, high := binarySearchBetween(cut, x, l)
	fmt.Println(cut[high] - cut[low])
}

func binarySearchBetween(cut []int, x, l int) (int, int) {
	low, high := 0, len(cut)-1
	for high-low != 1 {
		median := (low + high) / 2
		if cut[median] < x {
			low = median
		} else {
			high = median
		}
	}
	return low, high
}

func main() {
	sc.Split(bufio.ScanWords)
	l, q := nextInt(), nextInt()
	input := make([]Query, q)
	for i := range input {
		input[i] = Query{nextInt(), nextInt()}
	}
	cut := []int{0, l}
	for _, query := range input {
		if query.c == 1 {
			_, pos := binarySearchBetween(cut, query.x, l)
			cut = append(cut[:pos+1], cut[pos:]...)
			cut[pos] = query.x
		} else {
			print(cut, query.x, l)
		}
	}
}
