package main

import "fmt"

func main() {

	var a, b int
	var s string
	fmt.Scanln(&a, &s, &b)

	fmt.Println(a + b)
}
