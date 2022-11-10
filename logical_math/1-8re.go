package main

import "fmt"

func main() {
	//harf_adder
	var a, b bool
	fmt.Scan(&a, &b)
	c := a && b
	s := (a || b) && !(a && b)
	if c {
		fmt.Printf("1 ")
	} else {
		fmt.Printf("0 ")
	}
	if s {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}

}
