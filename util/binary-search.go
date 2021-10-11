package main

import "fmt"

// 要素が存在するかどうかをチェック
func binarySearch(target int, array []int) bool {
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2
		if array[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(array) || array[low] != target {
		return false
	}
	return true
}

func binaryMedLen(target int, array []int) int {
	low := 0
	high := len(array) - 1
	for high-low != 1 {
		median := (low + high) / 2
		if array[median] < target {
			low = median
		} else {
			high = median
		}
	}
	return array[high] - array[low]
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
