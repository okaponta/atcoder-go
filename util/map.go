package util

import "fmt"

func init() {
	contestMap := map[string]int{"key1": 0, "key2": 1}
	contestMap = make(map[string]int)

	// set的な
	set := make(map[int]struct{})
	set[0] = struct{}{}

	for key, value := range contestMap {
		fmt.Println(key, value)
	}
}
