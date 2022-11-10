package main

import "fmt"

func main() {

	var n, k int
	fmt.Scan(&n, &k)
	var m int
	for {
		n *= 2

		if n >= k {
			fmt.Println(m)
			break
		}
		m++
	}
}
