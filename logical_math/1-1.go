package main

import "fmt"

func main() {

	var a, b bool
	fmt.Scan(&a, &b)
	if a && b == true {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
