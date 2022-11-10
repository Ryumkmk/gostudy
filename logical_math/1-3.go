package main

import "fmt"

func main() {
	var a bool
	fmt.Scan(&a)
	if a == true {
		//not a
		fmt.Println("0")
	} else {
		fmt.Println("1")
	}

}
