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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)
	n, x, y := nextInt(), nextInt(), nextInt()

	count := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		count[i] = make([][]int, x+1)
		for j := 0; j <= x; j++ {
			count[i][j] = make([]int, y+1)
			for k := 0; k <= y; k++ {
				count[i][j][k] = 1000
			}
		}
	}
	count[0][0][0] = 0
	for i := 1; i <= n; i++ {
		a, b := nextInt(), nextInt()
		for j := 0; j <= x; j++ {
			for k := 0; k <= y; k++ {
				count[i][j][k] = min(count[i][j][k], count[i-1][j][k])
				if count[i-1][j][k] == 1000 {
					continue
				}
				jj := min(j+a, x)
				kk := min(k+b, y)
				count[i][jj][kk] = min(count[i][jj][kk], count[i-1][j][k]+1)
			}
		}
	}
	if count[n][x][y] == 1000 {
		count[n][x][y] = -1
	}
	fmt.Println(count[n][x][y])
}
