package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s = sc.Text()
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("^")
	}
}
