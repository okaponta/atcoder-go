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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	player := make(map[int]int, n)
	for i := range a {
		a[i] = nextInt()
		player[a[i]] = i + 1
	}
	sort.Ints(a)
	fmt.Println(player[a[len(a)-2]])
}
