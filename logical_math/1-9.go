package main

import "fmt"

func harfadder(a bool, b bool) (bool, bool) {
	//harf_adder + full_adder

	fmt.Scan(&a, &b)
	cx := a && b
	sy := (a || b) && !(a && b)
	return cx, sy
}

func xor(a, b bool) bool {
	cx := (a || b) && !(a && b)
	return cx

}

func main() {

	var a, b, cx, sy, cy, c1, c2, s bool
	fmt.Scan(&a, &b, &c1)

	cx, sy = harfadder(a, b)
	cy, s = harfadder(sy, c1)
	c2 = xor(cx, cy)
	if c2 {
		fmt.Printf("1 ")
	} else {
		fmt.Printf("0 ")
	}
	if s {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}

}
