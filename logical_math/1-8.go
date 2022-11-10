package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)

	c := a + b
	var out = make([]int, 0)
	if c == 1 {
		out = append(out, 1, 0)
	}
	if c == 0 {
		fmt.Println("0 0")
	} else {
		for c != 1 {
			out = append(out, c%2)
			c /= 2
			if c == 1 {
				out = append(out, 1)
			}
		}
	}

	for i := len(out); i > 0; i-- {
		fmt.Printf("%d", out[i-1])
		if i != 1 {

			fmt.Printf(" ")
		}
	}
}
