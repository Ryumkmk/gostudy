package main

import "fmt"

func main() {
	//XNOR
	var a, b bool
	fmt.Scan(&a, &b)

	if !((a || b) && (!(a && b))) == true {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
