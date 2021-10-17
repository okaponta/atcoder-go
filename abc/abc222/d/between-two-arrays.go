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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a, b := make([]int, n), make([]int, n)
	mod := 998244353
	for i := range a {
		a[i] = nextInt()
	}
	for i := range b {
		b[i] = nextInt()
	}
	initMap := make(map[int]int)
	for i := a[0]; i <= b[0]; i++ {
		initMap[i] = i - a[0] + 1
	}
	before := initMap
	for i := 1; i < n; i++ {
		after := make(map[int]int)
		_, ok := before[a[i]]
		if ok {
			after[a[i]] = before[a[i]]
		} else {
			after[a[i]] = before[b[i-1]]
		}
		for j := a[i] + 1; j <= b[i]; j++ {
			val, ok := before[j]
			if !ok {
				val = before[b[i-1]]
			}
			after[j] = (after[j-1] + val) % mod
		}
		before = after
	}
	fmt.Println(before[b[n-1]])
}
