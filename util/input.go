package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n string
	fmt.Scanf("%s", &n)

	var i int
	fmt.Scanf("%d", &i)
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
