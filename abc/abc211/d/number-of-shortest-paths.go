package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

var sc = bufio.NewScanner(os.Stdin)
var paths map[int][]int
var towns []town
var q = make([]int, 0)
var qmap = make(map[int]struct{})

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type town struct {
	time, way int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	towns = make([]town, n)
	paths = make(map[int][]int)
	for i := range towns {
		towns[i] = town{mod, 0}
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		patha, oka := paths[a]
		if oka {
			paths[a] = append(patha, b)
		} else {
			paths[a] = []int{b}
		}
		pathb, okb := paths[b]
		if okb {
			paths[b] = append(pathb, a)
		} else {
			paths[b] = []int{a}
		}
	}

	towns[0].time = 0
	towns[0].way = 1

	// 要素をいれる
	q = append(q, 0)
	qmap[0] = struct{}{}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		delete(qmap, v)
		nextTown(v)
	}
	fmt.Println(towns[n-1].way)
}

func nextTown(i int) {
	for _, next := range paths[i] {
		if towns[next].time > towns[i].time+1 {
			towns[next].time = towns[i].time + 1
			towns[next].way = towns[i].way
		} else if towns[next].time == towns[i].time+1 {
			towns[next].way = (towns[next].way + towns[i].way) % mod
		} else {
			continue
		}
	}
	for _, next := range paths[i] {
		if towns[next].time >= towns[i].time+1 {
			_, ok := qmap[next]
			if !ok {
				q = append(q, next)
				qmap[next] = struct{}{}
			}
		}
	}
}
