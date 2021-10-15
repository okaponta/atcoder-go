package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	forecast := []byte(s)
	for i := range forecast {
		if i+1 == n {
			if forecast[i] == 'o' {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
}
