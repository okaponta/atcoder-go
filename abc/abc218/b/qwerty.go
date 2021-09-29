package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)

	var line string
	if sc.Scan() {
		line = sc.Text()
	}

	answer := ""

	for _, p := range strings.Split(line, " ") {
		i, _ := strconv.Atoi(p)
		answer += string(rune('a' - 1 + i))
	}

	fmt.Println(answer)
}
