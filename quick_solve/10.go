package main

import "fmt"

func main() {

	var n int
	var sharp string = "##########"
	var point string = ".........."
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		fmt.Println(sharp)
		fmt.Println(point)

	}
}
