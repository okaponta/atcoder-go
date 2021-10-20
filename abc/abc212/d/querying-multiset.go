package main

import (
	"bufio"
	"container/heap"
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
	defer wr.Flush()
	q := nextInt()
	var offset int64
	pq := priorityQueue{}
	for i := 0; i < q; i++ {
		p := nextInt()
		if p == 1 {
			x := int64(nextInt()) - offset
			heap.Push(&pq, pqi{x})
		} else if p == 2 {
			x := nextInt()
			offset += int64(x)
		} else {
			if len(pq) == 0 {
				continue
			}
			a := pq[0].a
			heap.Pop(&pq)
			fmt.Fprintln(wr, a+offset)
		}

	}
}

type pqi struct{ a int64 }

type priorityQueue []pqi

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].a < pq[j].a }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqi)) }
func (pq *priorityQueue) Pop() interface{} {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}
