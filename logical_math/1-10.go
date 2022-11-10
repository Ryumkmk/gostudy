package main

import "fmt"

func xor(a, b bool) bool {
	c := (a || b) && !(a && b)
	return c

}

func and(a, b bool) bool {
	c := a && b
	return c
}

func or(a, b bool) bool {
	c := a || b
	return c
}

func not(a bool) bool {
	return !a
}

func main() {
	var a, b, c, d bool
	fmt.Scan(&a, &b, &c, &d)
	out := xor(not(or(and(not(a), not(b)), not(c))), d)
	if out {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
