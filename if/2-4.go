package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	if x >= 10 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
