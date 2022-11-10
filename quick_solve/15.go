package main

import "fmt"

func main() {
	var t1, t2 int
	fmt.Scan(&t1, &t2)
	if (t1 > t2) == true {
		fmt.Printf("%d", t2-t1)
	} else if t1 < t2 {
		fmt.Printf("+%d", t2-t1)
	} else {
		fmt.Println("0")
	}
}
