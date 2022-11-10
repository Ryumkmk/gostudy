package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	digit := strconv.Itoa(n)
	// fmt.Println(utf8.RuneCountInString(digit))
	fmt.Println(len(digit))
}
