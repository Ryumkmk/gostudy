package main

import "fmt"

func add(n, x int) int {
	return n + x
}

func sub(n, x int) int {
	return n - x
}

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	if (add(add(n, a), b) == 0) || (add(sub(n, a), b) == 0) {
		fmt.Println("YES")
	} else if (sub(add(n, a), b) == 0) || (sub(sub(n, a), b) == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
