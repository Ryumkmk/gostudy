package main

import "fmt"

func main() {

	var n, k, t int
	fmt.Scan(&n, &k, &t)

	cliff := float64(t) + 0.1
	if float64(n*k) > cliff {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
