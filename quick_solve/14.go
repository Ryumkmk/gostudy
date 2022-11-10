package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i <= 9; i++ {
		fmt.Println(i)
	}
	for i := 0; i < n; i++ {
		fmt.Println(i)

	}
}
