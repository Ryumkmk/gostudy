package main

import "fmt"

func main() {

	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Printf("%d %d", a*a, b*b+c*c)
}
