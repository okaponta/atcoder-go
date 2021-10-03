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

func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if a[0] == e[0] && a[1] == e[1] {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	coordinate := make([][]int, n)

	for i := range coordinate {
		coordinate[i] = []int{nextInt(), nextInt()}
	}

	sort.Slice(coordinate, func(i, j int) bool {
		return coordinate[i][0] < coordinate[j][0] && coordinate[i][1] < coordinate[j][1]
	})

	answer := 0

	for i := range coordinate {
		for j := range coordinate {
			if !(coordinate[i][0] < coordinate[j][0] && coordinate[i][1] < coordinate[j][1]) {
				// もし1点目より2点目の方が右上になかったら処理終了
				continue
			}
			if contains(coordinate, []int{coordinate[i][0], coordinate[j][1]}) &&
				contains(coordinate, []int{coordinate[j][0], coordinate[i][1]}) {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
