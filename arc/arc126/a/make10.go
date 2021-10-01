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

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func calc10(n2, n3, n4 int64) {
	var answer int64
	// 3,3,4
	case1 := min(n3/2, n4)
	answer += case1
	n3 -= case1 * 2
	n4 -= case1
	// 2,2,3,3
	case2 := min(n2/2, n3/2)
	answer += case2
	n2 -= case2 * 2
	n3 -= case2 * 2
	// 2,4,4
	case3 := min(n2, n4/2)
	answer += case3
	n2 -= case3
	n4 -= case3 * 2
	// 2,2,2,4
	case4 := min(n2/3, n4)
	answer += case4
	n2 -= case4 * 3
	n4 -= case4
	// 2,2,2,2,2
	case5 := n2 / 5
	answer += case5
	fmt.Println(answer)
}

func main() {
	sc.Split(bufio.ScanWords)
	t := nextInt()

	testCase := make([][]int64, t)

	for i := range testCase {
		testCase[i] = []int64{nextInt64(), nextInt64(), nextInt64()}
	}
	for _, j := range testCase {
		calc10(j[0], j[1], j[2])
	}
}
