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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	candy := make([]int, n)
	dist := make(map[int]int)
	ans := 0
	for i := range candy {
		candy[i] = nextInt()
		_, ok := dist[candy[i]]
		if ok {
			dist[candy[i]]++
		} else {
			dist[candy[i]] = 1
		}
		if i >= k-1 {
			ans = max(ans, len(dist))
			if dist[candy[i-k+1]] == 1 {
				delete(dist, candy[i-k+1])
			} else {
				dist[candy[i-k+1]]--
			}
		}
	}
	fmt.Println(ans)
}
