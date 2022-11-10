package main

import "fmt"

func main() {

	var f1, f2 int
	fmt.Scan(&f1, &f2)
	if (f1 > f2) == true {
		fmt.Println(f1 - f2)
	} else {
		fmt.Println(f2 - f1)
	}

}
