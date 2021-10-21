package main

import (
	"bufio"
	"container/heap"
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
	n, m := nextInt(), nextInt()
	input := make([][]int, n)
	count := make([]int, n)
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		input[a] = append(input[a], b)
		count[b]++
	}
	pq := priorityQueue{}

	for b, cnt := range count {
		if cnt == 0 {
			// bがないやつ
			pq.Push(pqi{b})
		}
	}

	ans := []string{}
	for pq.Len() > 0 {
		// 小さい順に要素取り出し
		a := pq[0].a
		heap.Pop(&pq)
		ans = append(ans, strconv.Itoa(a+1))
		for _, b := range input[a] {
			count[b]--
			if count[b] == 0 {
				// bがなくなったので候補に追加
				heap.Push(&pq, pqi{b})
			}
		}
	}

	if len(ans) < n {
		// ループがあるということ
		fmt.Println(-1)
		return
	}

	joined := strings.Join(ans, " ")
	fmt.Println(joined)
}

type pqi struct{ a int }

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
