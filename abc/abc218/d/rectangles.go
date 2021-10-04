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

type Coordinate struct {
	x, y int
}

type Range struct {
	low, high int
}

func compare(a, b Coordinate) bool {
	if a.x == b.x {
		return a.y < b.y
	}
	return a.x < b.x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	list := make([]Coordinate, n)
	for i := range list {
		list[i] = Coordinate{nextInt(), nextInt()}
	}

	// xが同じのyの最低値と最高値の数
	pair := make(map[Coordinate]int)

	for i := range list {
		for j := i + 1; j < n; j++ {
			if list[i].x == list[j].x {
				low := list[i].y
				high := list[j].y
				if low > high {
					low, high = high, low
				}
				pair[Coordinate{low, high}]++
			}
		}
	}

	answer := 0

	for _, n := range pair {
		// 1+2+3...なので
		answer += n * (n - 1) / 2
	}

	fmt.Println(answer)
}
