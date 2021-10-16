package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextBytes() []byte {
	sc.Scan()
	return []byte(sc.Text())
}

type Player struct {
	num, win int
}

func jj(a, b byte) int {
	if a == b {
		return 0
	}
	if a == 'G' && b == 'C' || a == 'C' && b == 'P' || a == 'P' && b == 'G' {
		return 1
	}
	return -1
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m := nextInt(), nextInt()
	janken := make([][]byte, 2*n)
	for i := range janken {
		janken[i] = nextBytes()
	}
	rank := make([]Player, 2*n)
	for i := range rank {
		rank[i] = Player{i, 0}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			match := jj(janken[rank[2*j].num][i], janken[rank[2*j+1].num][i])
			if match == 1 {
				rank[2*j].win++
			} else if match == -1 {
				rank[2*j+1].win++
			}
		}
		sort.Slice(rank, func(i, j int) bool {
			if rank[i].win == rank[j].win {
				return rank[i].num < rank[j].num
			}
			return rank[i].win > rank[j].win
		})
	}
	for _, p := range rank {
		fmt.Fprintln(wr, p.num+1)
	}
}
