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

func f(x, y int) int {
	return (x + y) % 10
}

func g(x, y int) int {
	return (x * y) % 10
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	pslice := make([]int, n)
	for i := range pslice {
		pslice[i] = nextInt()
	}

	var arr [10]int

	for i, p := range pslice {
		if i == 0 {
			arr[p] = 1
			continue
		}
		var newarr [10]int
		for ii, jj := range arr {
			newarr[f(ii, p)] += jj
			newarr[f(ii, p)] %= 998244353
			newarr[g(ii, p)] += jj
			newarr[g(ii, p)] %= 998244353
		}
		arr = newarr
	}

	for _, j := range arr {
		fmt.Println(j)
	}
}
