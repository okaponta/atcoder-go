package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	var k int
	var a string
	fmt.Scanf("%d", &k)
	fmt.Scanf("%s", &a)
	defer wr.Flush()
	fmt.Fprintln(wr, "hoge")
}
