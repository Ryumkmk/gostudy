package main

import "fmt"

func main() {
	var a, b, c int
	n := 0
	fmt.Scan(&a, &b, &c)
	n += a
	n *= b
	n %= c
	fmt.Println(n)
}
