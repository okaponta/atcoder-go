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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func binaryMin(target int, array []int) int {
	low := 0
	high := len(array) - 1
	for high-low != 1 {
		median := (low + high) / 2
		if array[median] < target {
			low = median
		} else {
			high = median
		}
	}
	return min(array[high]-target, target-array[low])
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := nextInt(), nextInt()
	n, m := make([]int, N), make([]int, M)
	for i := range n {
		n[i] = nextInt()
	}
	for i := range m {
		m[i] = nextInt()
	}
	sort.Ints(n)
	nmin, nmax := n[0], n[len(n)-1]
	ans := 2147483647
	for _, mi := range m {
		if mi == nmin || mi == nmax || ans == 0 {
			ans = 0
			break
		}
		if mi < nmin {
			ans = min(ans, nmin-mi)
			continue
		}
		if nmax < mi {
			ans = min(ans, mi-nmax)
			continue
		}
		ans = min(ans, binaryMin(mi, n))
	}
	fmt.Println(ans)
}
