package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	contestMap := map[string]int{"ABC": 0, "ARC": 1, "AGC": 2, "AHC": 3}
	valueMap := map[int]string{0: "ABC", 1: "ARC", 2: "AGC", 3: "AHC"}
	val := 6 - (contestMap[nextStr()] + contestMap[nextStr()] + contestMap[nextStr()])
	fmt.Println(valueMap[val])
}
