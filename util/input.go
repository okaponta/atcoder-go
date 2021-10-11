package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readString() string {
	var n string
	fmt.Scanf("%s", &n)
	return n
}

func readInt() int {
	var i int
	fmt.Scanf("%d", &i)

	var n int
	fmt.Scan(&n) // ちょっと遅いかも

	return i
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	nextInt()
}
