package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a >= 10 && b >= 10 {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}

}
