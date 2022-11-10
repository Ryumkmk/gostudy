package main

import "fmt"

func main() {
	var f int
	fmt.Scan(&f)
	if f >= 20 && f <= 15000 {
		fmt.Println("yes")
	} else if f > 1500 && f <= 20000 {
		fmt.Println("not sure")
	} else {
		fmt.Println("no")
	}
}
