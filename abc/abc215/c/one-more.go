package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

func runeToInt(before []rune) []int {
	converted := []int{}
	for _, r := range before {
		converted = append(converted, int(r-'a'))
	}
	return converted
}

func IntToRune(before []int) []rune {
	converted := []rune{}
	for _, i := range before {
		converted = append(converted, rune(i+'a'))
	}
	return converted
}

func next(before []int) []int {
	next_permutation(before)
	return before
}

func main() {
	sc.Split(bufio.ScanWords)
	s, k := nextStr(), nextInt()
	run := []rune(s)

	sort.Slice(run, func(i, j int) bool { return run[i] < run[j] })
	slice := runeToInt(run)
	for i := 1; i < k; i++ {
		slice = next(slice)
	}
	run = IntToRune(slice)
	fmt.Println(string(run))
}

func next_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return x < y })
}

func prev_permutation(a []int) bool {
	return _permutation_pandita(a, func(x, y int) bool { return y < x })
}

func _permutation_pandita(a []int, less func(x, y int) bool) bool {
	i := len(a) - 2
	// Find the right most incresing elements a[i] and a[i+1]
	for 0 <= i && !less(a[i], a[i+1]) {
		i--
	}
	if i == -1 {
		return false
	}
	j := i + 1
	// Find the smallest element that is greater than a[i] in a[i+1:]
	for k := j + 1; k < len(a); k++ {
		// a[i] < a[k] && a[k] <= a[j]
		if less(a[i], a[k]) && !less(a[j], a[k]) {
			j = k
		}
	}
	a[i], a[j] = a[j], a[i]
	for p, q := i+1, len(a)-1; p < q; p, q = p+1, q-1 {
		a[p], a[q] = a[q], a[p]
	}
	return true
}
