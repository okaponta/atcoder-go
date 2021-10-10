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

type UserChange struct {
	days  int
	count int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	userChanges := make(map[int]int)
	for i := 0; i < n; i++ {
		startDate := nextInt()
		endDate := startDate + nextInt()
		userChanges[startDate] += 1
		userChanges[endDate] -= 1
	}

	for key, value := range userChanges {
		if value == 0 {
			delete(userChanges, key)
		}
	}

	userSlice := make([]UserChange, len(userChanges))
	count := 0
	for key, value := range userChanges {
		userSlice[count] = UserChange{key, value}
		count++
	}

	sort.Slice(userSlice, func(i, j int) bool { return userSlice[i].days < userSlice[j].days })

	answer := make([]int, n+1)

	beforeDay := 0
	countAns := 0
	for i, userChange := range userSlice {
		if i == 0 {
			beforeDay = userChange.days
			countAns = userChange.count
			continue
		}
		answer[countAns] += userChange.days - beforeDay
		beforeDay = userChange.days
		countAns += userChange.count
	}

	for i, j := range answer {
		if i == 0 {
			continue
		}
		if i != 1 {
			fmt.Print(" ")
		}
		fmt.Print(j)
	}
	fmt.Println()
}
