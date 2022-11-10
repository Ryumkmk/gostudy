package main

import (
	"fmt"
	"math"
)

func main() {
	a := 202
	b := 134
	c := 107
	fmt.Println(int(math.Pow(float64(((a + b) * c)), 2)))
}
