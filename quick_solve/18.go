package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	if s == "Friday" {
		fmt.Println("TGIF")
	} else {
		fmt.Println("Still " + s)
	}
}
