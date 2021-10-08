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
	n, m := nextInt(), nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}

	primeMap := make(map[int]struct{})
	for _, ai := range a {
		if ai == 1 {
			continue
		}
		num := 2
		for num*num <= ai {
			if ai%num == 0 {
				ai /= num
				primeMap[num] = struct{}{}
			} else {
				num++
			}
		}
		primeMap[ai] = struct{}{}
	}

	notAnswer := make([]bool, m+1)
	notAnswer[0] = true
	for i := range primeMap {
		for j := i; j <= m; j += i {
			notAnswer[j] = true
		}
	}

	ans := make([]int, 0)
	for i, ng := range notAnswer {
		if !ng {
			ans = append(ans, i)
		}
	}
	fmt.Println(len(ans))
	for i := range ans {
		fmt.Println(ans[i])
	}
}
