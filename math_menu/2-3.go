package main

import "fmt"

func main() {

	n := 10000
	n /= 361
	n %= 28
	fmt.Println(n)
}
