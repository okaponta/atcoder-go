package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	q := make([]string, nextInt())
	for i := range q {
		q[nextInt()-1] = strconv.Itoa(i + 1)
	}
	fmt.Println(strings.Join(q, " "))
}
