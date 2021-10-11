package util

import (
	"fmt"
	"sort"
)

func init() {
	slice := []int{0, 1}
	slice = make([]int, 2)

	for i, j := range slice {
		fmt.Println(i, j)
	}
}

func sortSlice() {
	slice := []int{0, 1}
	sort.Ints(slice)
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	// カスタム
	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })
}
