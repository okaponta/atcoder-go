package util

import (
	"container/heap"
	"fmt"
)

func main() {
	pq := priorityQueue{}
	// 要素をいれる
	heap.Push(&pq, pqi{1})
	// 最初の要素を取り出す
	a := heap.Pop(&pq)
	fmt.Println(a)
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
