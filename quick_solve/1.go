package main

import "fmt"

func main() {

	var d, s int
	fmt.Scan(&d, &s)
	if d*100000/s >= 10000 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
