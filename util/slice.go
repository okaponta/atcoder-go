package util

import (
	"fmt"
	"sort"
)

func init() {
	slice := []int{0, 1}
	slice = make([]int, 2)

	sort.Slice(slice, func(i, j int) bool { return slice[i] < slice[j] })

	for i, j := range slice {
		fmt.Println(i, j)
	}
}
