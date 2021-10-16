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
	if a == 'G' {
		if b == 'C' {
			return 1
		}
		if b == 'P' {
			return -1
		}
		return 0
	}
	if a == 'C' {
		if b == 'P' {
			return 1
		}
		if b == 'G' {
			return -1
		}
		return 0
	}
	if a == 'P' {
		if b == 'G' {
			return 1
		}
		if b == 'C' {
			return -1
		}
		return 0
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()
	n, m := nextInt(), nextInt()
	playerNum := n * 2
	janken := make([][]byte, playerNum)
	for i := range janken {
		janken[i] = nextBytes()
	}
	leaderBoard := make([]Player, playerNum)
	for i := range leaderBoard {
		leaderBoard[i] = Player{i, 0}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a := leaderBoard[2*j]
			b := leaderBoard[2*j+1]
			aj := janken[a.num][i]
			bj := janken[b.num][i]
			match := jj(aj, bj)
			if match == 1 {
				leaderBoard[2*j].win++
			} else if match == -1 {
				leaderBoard[2*j+1].win++
			}
		}
		sort.Slice(leaderBoard, func(i, j int) bool {
			if leaderBoard[i].win == leaderBoard[j].win {
				return leaderBoard[i].num < leaderBoard[j].num
			}
			return leaderBoard[i].win > leaderBoard[j].win
		})
	}
	for _, p := range leaderBoard {
		fmt.Fprintln(wr, p.num+1)
	}
}
