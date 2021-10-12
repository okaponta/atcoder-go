package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func printBool(ok bool) {
	if ok {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func printSlice(slice []string) {
	joined := strings.Join(slice, " ")
	fmt.Println(joined)
}

var wr = bufio.NewWriter(os.Stdout)

func main() {
	defer wr.Flush()
	fmt.Fprintln(wr, "hoge")
}
