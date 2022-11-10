package main

import "fmt"

func main() {

	var s string
	fmt.Scan(&s)
	if s == "paiza" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
