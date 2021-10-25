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

type balls struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	top := make(map[int]int)
	nextMap := make(map[int]balls)
	for i := 0; i < m; i++ {
		k := nextInt()
		a := make([]int, k)
		for j := range a {
			a[j] = nextInt()
			if j == 0 {
				top[a[j]]++
				continue
			}
			bs, ok := nextMap[a[j-1]]
			if ok {
				bs.b = a[j]
				nextMap[a[j-1]] = bs
			} else {
				nextMap[a[j-1]] = balls{a[j], 0}
			}
		}
	}

	removable := make([]int, 0)

	for key, value := range top {
		if value == 2 {
			removable = append(removable, key)
		}
	}

	count := 0

	for len(removable) > 0 {
		remove := removable[0]
		removable[0] = removable[len(removable)-1]
		removable = removable[:len(removable)-1]
		delete(top, remove)
		a, b := nextMap[remove].a, nextMap[remove].b
		if a != 0 {
			top[a]++
			if top[a] == 2 {
				removable = append(removable, a)
			}
		}
		if b != 0 {
			top[b]++
			if top[b] == 2 {
				removable = append(removable, b)
			}
		}
		count++
	}

	if count == n {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
