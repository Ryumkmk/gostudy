package main

import "fmt"

func main() {

	var x, y, z int
	fmt.Scan(&x, &y, &z)
	if (x >= 10 && y >= 10) || z >= 10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
