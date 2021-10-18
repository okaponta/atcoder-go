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

type D struct {
	a, b int
	sec  float64
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	doukasen := make([]D, n)
	var sum float64
	for i := range doukasen {
		var sec float64
		a, b := nextInt(), nextInt()
		sec = float64(a) / float64(b)
		doukasen[i] = D{a, b, sec}
		sum += sec
	}
	len := 0
	remain := sum / 2
	var targetD int
	for i := range doukasen {
		if remain-doukasen[i].sec < 0 {
			targetD = i
			break
		}
		len += doukasen[i].a
		remain -= doukasen[i].sec
	}
	fmt.Println(float64(len) + float64(doukasen[targetD].b)*remain)
}
