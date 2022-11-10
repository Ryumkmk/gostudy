package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a_i int
		fmt.Scan(&a_i)
		if a_i == 0 {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
