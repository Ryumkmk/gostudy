package main

import "fmt"

func main() {
	var n, k, t int
	fmt.Scan(&n, &k, &t)
	if k*t%n == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
