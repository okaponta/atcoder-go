package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	answerStr := make([]string, n)

	for i := 1; i < len(answer); i++ {
		answerStr[i-1] = strconv.Itoa(answer[i])
	}

	fmt.Println(strings.Join(answerStr, " "))
}
