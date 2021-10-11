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

type Query struct {
	c, x int
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	l, q := nextInt(), nextInt()
	input := make([]Query, q)
	vals := make([]int, q+2)
	for i := range input {
		input[i] = Query{nextInt(), nextInt()}
		vals[i] = input[i].x
	}
	vals[q] = l
	vals[q+1] = 0
	sort.Ints(vals)

	valToIndexMap := map[int]int{}
	for i, v := range vals {
		valToIndexMap[v] = i
	}

	bit := newBinaryIndexedTree(q + 1)

	for _, query := range input {
		if query.c == 1 {
			bit.add(valToIndexMap[query.x], 1)
		} else {
			s := bit.sum(valToIndexMap[query.x])
			fmt.Fprintln(wr, vals[bit.lowerBound(s+1)]-vals[bit.lowerBound(s)])
		}
	}
}

type binaryIndexedTree []int

func newBinaryIndexedTree(n int) binaryIndexedTree {
	return make(binaryIndexedTree, n+1)
}

func (t binaryIndexedTree) add(i, x int) {
	for i++; i < len(t) && i > 0; i += i & -i {
		t[i] += x
	}
}

func (t binaryIndexedTree) sum(i int) int {
	s := 0
	for i++; i > 0; i -= i & -i {
		s += t[i]
	}
	return s
}

func (t binaryIndexedTree) lowerBound(x int) int {
	idx, k := 0, 1
	for k < len(t) {
		k <<= 1
	}
	for k >>= 1; k > 0; k >>= 1 {
		if idx+k < len(t) && t[idx+k] < x {
			x -= t[idx+k]
			idx += k
		}
	}
	return idx
}
