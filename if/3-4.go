package main

import (
	"fmt"
)

func main() {
	// weekday := time.Weekday(6)
	var weekday = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	var x int
	fmt.Scan(&x)
	i := x % 7
	if i == 0 {
		fmt.Println(weekday[6])
		return
	}
	fmt.Println(weekday[i-1])
}
