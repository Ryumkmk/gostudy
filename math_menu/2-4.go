package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	n := 10000
	n /= a
	n %= b
	fmt.Println(n)
}
