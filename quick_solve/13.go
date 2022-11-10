package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	fmt.Printf("%d"+"%%", int((float64(n)/float64(140))*100))

}
