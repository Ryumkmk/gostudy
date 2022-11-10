package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string
	var c = []string{"CEFGHIJKLMNSTUVWXYZ", "ADOPQR", "B"}
	fmt.Scan(&s)
	for i := 0; i < 3; i++ {
		if strings.Contains(c[i], s) == true {
			fmt.Println(i)
			break
		}
	}

}
