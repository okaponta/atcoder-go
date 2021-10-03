package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	existsMap := make(map[Coordinate]bool)

	for i := range list {
		list[i] = Coordinate{nextInt(), nextInt()}
		existsMap[list[i]] = true
	}

	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j])
	})

	answer := 0

	for i := range list {
		for j := i; j < n; j++ {
			if !(list[i].x < list[j].x && list[i].y < list[j].y) {
				// もし1点目より2点目の方が右上になかったら処理終了
				continue
			}
			if existsMap[Coordinate{list[i].x, list[j].y}] &&
				existsMap[Coordinate{list[j].x, list[i].y}] {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
