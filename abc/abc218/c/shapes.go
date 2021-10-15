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

func readBool(n int) []bool {
	sc.Scan()
	line := []byte(sc.Text())
	res := make([]bool, n)
	for i := range res {
		res[i] = line[i] == '#'
	}
	return res
}

func count(grid [][]bool) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				ans++
			}
		}
	}
	return ans
}

func find_left_top(grid [][]bool) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				return i, j
			}
		}
	}
	return 0, 0
}

func printBool(ok bool) {
	if ok {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func is_same(s, t [][]bool) bool {
	si, sj := find_left_top(s)
	ti, tj := find_left_top(t)
	offset_i, offset_j := ti-si, tj-sj
	for i := range s {
		for j := range s[i] {
			ii := i + offset_i
			jj := j + offset_j
			if 0 <= ii && ii < len(s) && 0 <= jj && jj < len(s) {
				if s[i][j] != t[ii][jj] {
					return false
				}
			} else if s[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := make([][]bool, n)
	t1, t2, t3, t4 := make([][]bool, n), make([][]bool, n), make([][]bool, n), make([][]bool, n)

	for i := range s {
		s[i] = readBool(n)
	}

	for i := range t1 {
		t1[i] = readBool(n)
		t2[i] = make([]bool, n)
		t3[i] = make([]bool, n)
		t4[i] = make([]bool, n)
	}

	if count(s) != count(t1) {
		fmt.Println("No")
		return
	}

	for i := range t1 {
		for j := range t1[i] {
			t2[j][n-i-1] = t1[i][j]
			t3[n-i-1][n-j-1] = t1[i][j]
			t4[n-j-1][i] = t1[i][j]
		}
	}
	ans := is_same(s, t1) || is_same(s, t2) || is_same(s, t3) || is_same(s, t4)
	printBool(ans)
}
