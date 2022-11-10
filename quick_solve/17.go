package main

import "fmt"

func main() {

	var l int
	fmt.Scan(&l)

	if l < 30 {
		fmt.Println("quiet")
	} else if l < 50 {
		fmt.Println("normal")
	} else if l < 70 {
		fmt.Println("noisy")
	} else {
		fmt.Println("very noisy")
	}
}
