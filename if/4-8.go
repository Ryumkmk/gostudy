package main

import "fmt"

func main() {

	var n, w int
	fmt.Scan(&n, &w)
	// n := 9342879
	// w := 10000000000
	// fmt.Println(n*w%4)
	if n%2 == 0 && w%2 == 0 {
		fmt.Println("YES")
		if n == 0 || w == 0 {
			fmt.Println("NO")
			return
		}
	} else {
		fmt.Println("NO")
	}
}
