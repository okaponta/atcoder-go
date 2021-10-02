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

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	a := make([]int, n)
	var sum int64

	for i := range a {
		a[i] = nextInt()
		sum += int64(a[i])
	}

	x := nextInt64()
	remain := x % sum

	count := (x / sum) * int64(n)

	for _, j := range a {
		count++
		remain -= int64(j)
		if remain < 0 {
			break
		}
	}

	fmt.Println(count)
}
