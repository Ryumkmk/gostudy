package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	b_mod := math.Pow((a+b)*c, 2)
	fmt.Println(math.Mod(b_mod, d))

}
