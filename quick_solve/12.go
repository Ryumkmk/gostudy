package main

import "fmt"

func main() {

	var s string
	var mt string = "Mt. "
	fmt.Scan(&s)
	b := fmt.Sprintf("%s%s", mt, s)
	fmt.Println(b)

}
